package service

import (
	"io"
	"math"
	"runtime"
	"strings"
	"sync"

	"github.com/yeka/zip"
)

const (
	TargetCharShort   = "0123456789abcdefghijklmnopqrstuvwxyz"
	TargetCharDefault = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	TargetCharLong    = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+-*/_@#$%"
	PassMaxLenDefault = 12
)

type SafeData struct {
	pass string
	done bool
	mux  sync.Mutex
}

func (f *SafeData) SetPass(pass string) {
	f.mux.Lock()
	f.pass = pass
	f.done = true
	f.mux.Unlock()
}

func (f *SafeData) SetDone(b bool) {
	f.mux.Lock()
	f.done = b
	f.mux.Unlock()
}

func (f *SafeData) CheckDone() bool {
	f.mux.Lock()
	defer f.mux.Unlock()
	return f.done
}

type SearchService struct {
	safeData   *SafeData
	targetChar string
	passMaxLen int
}

func NewSearchService(targetCharList string, targetCharType int, passMaxLen int) *SearchService {
	var t string
	if len(targetCharList) > 0 {
		t = targetCharList
	} else {
		switch targetCharType {
		case 1:
			t = TargetCharShort
		case 2:
			t = TargetCharLong
		default:
			t = TargetCharDefault
		}
	}
	p := PassMaxLenDefault
	if passMaxLen > 0 {
		p = passMaxLen
	}
	return &SearchService{
		safeData:   &SafeData{done: false},
		targetChar: t,
		passMaxLen: p,
	}
}

func (s *SearchService) SearchPassWord(input string) (string, error) {
	fullPath, err := GetFullPath(input)
	if err != nil {
		return "", err
	}
	if err := CheckFileOpen(fullPath); err != nil {
		return "", err
	}

	wg := new(sync.WaitGroup)
	for i := 0; i < runtime.NumCPU(); i++ {
		i := i
		wg.Add(1)
		go func(n int) {
			s.search(n, fullPath)
			wg.Done()
		}(i)
	}
	wg.Wait()

	return s.safeData.pass, nil
}

func (s *SearchService) search(offset int, fullPath string) {
	r, err := zip.OpenReader(fullPath)
	if err != nil {
		return
	}
	defer r.Close()

	for i := 0; i < s.passMaxLen; i++ {
		limit := int(math.Pow(float64(len(s.targetChar)), float64(i+1)))
		for j := offset; j < limit; j += runtime.NumCPU() {
			if s.safeData.CheckDone() {
				return
			}
			pass := s.getPassString(i, j)
			for _, f := range r.File {
				if f.IsEncrypted() {
					f.SetPassword(pass)
				}

				r, err := f.Open()
				if err != nil {
					continue
				}

				buf, err := io.ReadAll(r)
				if err != nil {
					continue
				}
				defer r.Close()

				if len(buf) <= 0 {
					continue
				}

				s.safeData.SetPass(pass)
				return
			}
		}
	}
}

func (s *SearchService) getPassString(i int, j int) string {
	res := make([]string, i+1)
	list := strings.Split(s.targetChar, "")
	for k := 0; k < i+1; k++ {
		div := int(math.Pow(float64(len(s.targetChar)), float64(k)))
		mod := int(math.Pow(float64(len(s.targetChar)), float64(k+1)))
		res[k] = list[int((j%mod)/div)]
	}
	return strings.Join(res, "")
}

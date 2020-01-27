package model

type Friends struct {
	Data []string
}

type Person struct {
	Name    string
	Friends *Friends
}

func (f *Friends) Add(name string) {
	f.Data = append(f.Data, name)
}

func (f *Friends) Remove(name string) {

}

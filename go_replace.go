package go_replace

type Replace struct{
	Id2Str map[int64]string
	Str2Id map[string]int64
}

func (s *Replace)AddReplace(str string)int64{
	if val, ok := s.Str2Id[str]; ok {
		return val
	}
	id := int64(len(s.Str2Id) + 1)
	s.Str2Id[str] = id
	s.Id2Str[id] = str
	return id
}

func (s *Replace)AddReplaceIndex(str string,index int64)int64{
	if val, ok := s.Str2Id[str]; ok {
		return val
	}
	s.Str2Id[str] = index
	s.Id2Str[index] = str
	return index
}

func NewReplace() *Replace{
	s := new(Replace)
	s.Str2Id = map[string]int64{}
	s.Id2Str = map[int64]string{}
	return s
}
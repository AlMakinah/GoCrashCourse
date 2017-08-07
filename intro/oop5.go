func (p *Person) ChangeName(newName string) (oldName string) {
	oldName = p.Name
	p.Name = newName
	return
}

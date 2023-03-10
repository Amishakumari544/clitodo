package todo 

import "time"

type item struct {
	Task string
	Done bool
	CreatedAt time.Time
	CompletedAt time.Time
}

type Todos []item


func (t *Todods) Add(task string){
	todo := time{
		Task : task,
		Done:false,
		CreatedAt : time.Now(),
		CompletedAt : time.Time{},
	}

	*t = append(*t, todo)


}


func (t *Todos) Compelete(index int) error{
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New(text : "Invalid index")
	}

	ls[index-1].CompletedAt = time.Now();
	ls[index-1].Done = true

	return nil

}

func (t *Todos) Delete(index int) error{
	ls := *t
	if index <= 0 || index > len(ls){
		return errors.New(text : "Invalid index")
	}

	*t = append(ls[:index-1], ls[index:]...)

	return nil
}

func (t *Todos) Load(filename string) error{
	file, err := ioutil.ReadFile(filename)
	if err != nil{
		if errors.Is(err, os.ErrNotExist){
			return nil
		}
		return err
	}
	if len(file) == 0{
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}
	return nil
}
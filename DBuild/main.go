package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/jcelliott/lumber"
	"golang.org/x/text/collate"
)

const version = "1.0.0"

type (
	Logger interface {
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}

	Driver struct {
		mutex   sync.Mutex
		mutexes map[string]*sync.Mutex
		dir     string
		log     Logger
	}
)

type Options struct {
	Logger
}

func New(dir string, options *Options) (*Driver, error) {
	dir = filepath.Clean(dir)

	opts := Options{}

	if options != nil {
		opts = *options
	}

	if opts.Logger == nil {
		opts.Logger = lumber.NewConsoleLogger((lumber.INFO))
	}

	driver := Driver{
		dir:     dir,
		mutexes: make(map[string]*sync.Mutex),
		log:     opts.Logger,
	}

	if _, err := os.Stat(dir); err != nil {
		opts.Logger.Debug("Using '%s' ", dir)
		return &driver, nil
	}

	opts.Logger.Debug("Creating the database at '%s'....", dir)
	return &driver, os.Mkdir(dir, 0755)

}

func stat(path string) (fi os.FileInfo, err error) {
	if fi, err = os.Stat(path); os.IsNotExist(err) {
		fi, err = os.Stat(path + ".json")
	}
	return
}

func (d *Driver) write(collection, resource string, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("Missing collection no place to save record")
	}

	if resource == "" {
		return fmt.Errorf("Missing resource - unable to save record")
	}

	mutex := d.getorCreateMitex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, collection)
	fnPath := filepath.Join(dir, resource+".json")
	tmpPath := fnPath + ".tmp"

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	b, err := json.MarshalIndent(v, "", "\t")

	if err != nil {
		return err
	}

	b = append(b, byte('\n'))

	if err := ioutil.WriteFile(tmpPath, b, 0644); err != nil {
		return err
	}

  return os.Rename(tmpPath, fnPath); 
}



func (d *Driver) Read(collection, resource string, v interface{}) error {

	if collection == "" {
		return fmt.Errorf("Missing collection no place to save record")
	}

	if resource == "" {
		return fmt.Errorf("Missing resource - unable to save record")
	}

   record := filepath.Join(d.dir , collection , resource);

  if _, err := stat(record); err != nil {
     return err;
  }

  b , err := ioutil.ReadFile(record + ".json");

  if err != nil {
     return err;
  }

  return json.Unmarshal(b , &v);

}

func (d* Driver) ReadAll(collection string)([]string , error) {
  if collection == "" {
      return nil, fmt.Errorf("Missing Collection");
  }

  dir := filepath.Join(d.dir , collection);

  if _ , err := stat(dir); err != nil {
     return nil , err;
  }
   
  files, _ := ioutil.ReadDir(dir);

  var records []string;

  for _ , file := range files {
    b , err := ioutil.ReadFile(filepath.Join(dir, file.Name()));

    if err != nil {
       return nil , err
    }

    records = append(records, string(b));
  }
  return records , nil;

}

func (d *Driver) Delet() error {

}

func (d *Driver) getorCreateMitex(collection string) *sync.Mutex {

	d.mutex.Lock()
	defer d.mutex.Unlock()
	m, ok := d.mutexes[collection]

	if !ok {
		m = &sync.Mutex{}
		d.mutexes[collection] = m
	}
	return m
}

type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func main() {
	dir := "./"

	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	employees := []User{
		{"Jhon", "23", "696969699", "Infosys", Address{"palampur", "hp", "india", "14352"}},
		{"Jogi", "26", "6964569699", "Infosysh", Address{"delhi", "it", "india", "14332"}},
		{"rock", "28", "696964799", "Infosysbb", Address{"up", "gp", "india", "14344"}},
	}

	for _, value := range employees {
		db.write("users", value.Name, User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(records)

	allusers := []User{}

	for _, f := range records {
		employeeFound := User{}
		if err := json.Unmarshal([]byte(f), &employeeFound); err != nil {
			fmt.Println("Errors", err)
		}
		allusers = append(allusers, employeeFound)
	}

	if err := db.Delet("user", "jhon"); err != nil {
		fmt.Println("Error", err)
	}

	if err := db.DeletAll("user", ""); err != nil {
		fmt.Println("Error", err)
	}

}

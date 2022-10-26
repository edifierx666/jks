package service

type Group struct {
  Jobs  *Jobs
  Cache *Cache
}

var Service *Group

func Init() {
  Service = &Group{
    Jobs:  NewJobs(),
    Cache: NewCache(),
  }
}

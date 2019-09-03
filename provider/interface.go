package provider


type ClusterProvider interface {
 Create(string) error
 Delete(string) error
 List() []string
}
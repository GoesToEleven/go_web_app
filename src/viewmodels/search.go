package viewmodels

func GetSearch() ViewModel {
    result := ViewModel{
        Title: "Search Stories",
        SignedIn: false,
        Active: "search",
    }
    return result
}
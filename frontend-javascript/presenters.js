/* Presenter */

//Presenters should only access public API of Router and Interactor.
//Here, we'll mark public API as capitalized name.

function NewLoginPagePresenter(router, interactor){
    function triggerLogin (username, password) {
        const response = interactor.Login(username, password)
        if (!response.error){
            const user = response.user
            router.GoToHomePage(user)
        }else{
            alert(response.error)
        }
    }

    //public APIs that can be accessed from View Controllers.
    //let's mark them as public by capitalizing the object property.
    const presenterObject = {
        LoginButtonClicked: triggerLogin
    }

    return presenterObject
}

function NewHomePagePresenter(router, interactor){
    function getAllExistingTodoList(){
        const response = interactor.GetAllExistingTodoList()
        if (!response.error){
            const allToDoList = response.allToDoList
            return allToDoList
        }else{
            alert(response.error)
            return []
        }
    }

    function triggerAddEntry(user, todoList, entryText){
        const response = interactor.AddNewToDoListEntry(user,todoList,entryText)
        if (!response.error){
            router.GoToHomePage(user) //refresh home page
        }else{
            alert(response.error)
        }
    }
    //public APIs that can be accessed from View Controllers.
    //let's mark them as public by capitalizing the object property.
    const presenterObject = {
        GetAllExistingTodoList: getAllExistingTodoList,
        AddEntryButtonClicked: triggerAddEntry,
    }

    return presenterObject
}
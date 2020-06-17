/* View Controllers */

//View Controllers should only access public API of Presenter.
//Here, we'll mark public API as capitalized name.

function NewLoginViewController(loginPagePresenter){
    function loadViewController(){
        removeAllChildrenFrom(rootView)
    
        const usernameTextBox = createTextBoxInline("Username", "text")
        const passwordTextBox = createTextBoxInline("Password", "password")
        const loginButton = createLink("Login", function(){
            const username = usernameTextBox.GetValue()
            const password = passwordTextBox.GetValue()
            loginPagePresenter.LoginButtonClicked(username, password)
        })
    
        rootView.append(usernameTextBox)
        rootView.append(passwordTextBox)
        rootView.append(loginButton)
    }

    //public APIs that can be accessed from Router.
    //let's mark them as public by capitalizing the object property.
    var viewController = {
        Load: loadViewController
    }
    return viewController
}

function NewHomeViewController(homePagePresenter){
    function loadViewController(user){
        removeAllChildrenFrom(rootView)
    
        const titleLabel = createLabelBlock(viewModelForHomePageTitle(user))
        rootView.append(titleLabel)

        const allToDoList = homePagePresenter.GetAllExistingTodoList()
        allToDoList.forEach(function(todoList){
            rootView.append(createToDoListView(todoList, function(todoList, entryText){
                console.log(`submitting new entry ${entryText} to ${todoList.Title} to-do list`)
                homePagePresenter.AddEntryButtonClicked(user,todoList,entryText)
            }))
        })
    }

    //public APIs that can be accessed from Router.
    //let's mark them as public by capitalizing the object property.
    var viewController = {
        Load: loadViewController
    }
    return viewController
}
/* Interactor */

function CreateInteractor(){
    function login(username, password){
        const path = "/login"
        const resp = callBackendMicroservice(path, {"username":username,"password":password})
        const response = {
            error: resp.error,
            user: resp.data?.user
        }
        return response
    }

    function getAllExistingTodoList(){
        const path = "/todo-lists"
        const resp = callBackendMicroservice(path)
        const response = {
            error: resp.error,
            allToDoList: resp.data?.allToDoList
        }
        return response
    }

    function addNewToDoListEntry(user,todoList,entryText) {
        const path = "/todo-list/add"
        const resp = callBackendMicroservice(path, {user,todoList,entryText})
        const response = {
            error: resp.error,
        }
        return response
    }

    //public APIs that can be accessed from Presenters.
    //let's mark them as public by capitalizing the object property.
    const interactor = {
        Login: login,
        GetAllExistingTodoList: getAllExistingTodoList,
        AddNewToDoListEntry: addNewToDoListEntry,
    }

    return interactor
}
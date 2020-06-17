/* Views */

//Views should not know anything about view controllers
//Don't pass in view controller into view like this: function createView(viewController){}

function createToDoListView(todoList, submitNewEntryHandler){
    var view = document.createElement("div");
    view.style = "margin-top: 10px;";
    view.append(createLabelBlock(`${todoList.Title} to-do list`))
    todoList.Entries.forEach(function (entry, index) {
        view.append(createToDoEntryView(entry))
    });
    view.append(createAddTodoListEntryView(todoList, submitNewEntryHandler))

    return view
}

function createAddTodoListEntryView(todoList, submitNewEntryHandler){
    var view = document.createElement("div");
    const newToDoEntryTextBox = createTextBoxInline("Add To-Do", "text")
    const addEntryButton = createLink("Add", function () {
        const entryText = newToDoEntryTextBox.GetValue()
        submitNewEntryHandler(todoList, entryText)
    })
    view.append(newToDoEntryTextBox)
    view.append(addEntryButton)

    return view
}

function createToDoEntryView(entry){
    var view = document.createElement("div");
    const entryIdLabel = createLabelInline("id:"+ entry.UUID.substring(0,4) + " | ")
    const entryOwnerLabel = createLabelInline("owner:"+ entry.Owner.Name + " | ")
    const entryTextLabel = createLabelInline(entry.Text)
    view.append(entryIdLabel)
    view.append(entryOwnerLabel)
    view.append(entryTextLabel)

    return view
}

function createTextBoxInline(title,type){
    var view = document.createElement("span");
    return createTextBox(view,title,type)
}

function createTextBoxBlock(title,type){
    var view = document.createElement("div");
    return createTextBox(view,title,type)
}

function createTextBox(view,title,type){
    var i = document.createElement("input"); //input element, text
    i.placeholder = title
    i.setAttribute('type',type);
    i.setAttribute('name',title);
    view.append(i)

    //public APIs that can be accessed from View Controllers.
    //let's mark them as public by capitalizing the object property.
    view.GetValue = function(){
        return i.value
    }
    return view
}

function createLabelInline(text){
    var view = document.createElement("span");
    return createLabel(view, text)
}

function createLabelBlock(text){
    var view = document.createElement("div");
    return createLabel(view, text)
}

function createLabel(view, text){
    var linkText = document.createTextNode(text);
    view.appendChild(linkText);

    return view
}

function createLink(text, clickEventHandler) {
    var view = document.createElement("span");
    var linkText = document.createTextNode(text);
    view.appendChild(linkText);
    view.style = "color:blue; text-decoration: underline; cursor: pointer;";
    //a.href = "#"
    view.onclick = clickEventHandler

    return view
}
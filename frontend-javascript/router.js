/* Router */

//global scope
function CreateRouter(){
    var availablePages = {}
    function injectAvailablePages(homePage, loginPage){
        availablePages.homePage = homePage
        availablePages.loginPage = loginPage
    }

    var activePageStack = []
    function goToLoginPage(){
        const loginViewController = availablePages.loginPage
        activePageStack.push(loginViewController)
        loginViewController.Load()
    }
    function goToHomePage(user){
        const homeViewController = availablePages.homePage
        activePageStack = [homeViewController] //empty the page hierarchy and add one page to it
        homeViewController.Load(user)
    }

    //public APIs that can be accessed from Presenters.
    //let's mark them as public by capitalizing the object property.
    const router = {
        Initialize: injectAvailablePages,
        GoToHomePage: goToHomePage,
        GoToLoginPage: goToLoginPage,
    }

    return router
}

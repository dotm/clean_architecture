console.log(pingBackendServer())

/* Setup Frontend */
const interactor = CreateInteractor()
const router = CreateRouter()

const loginPage = NewLoginViewController(NewLoginPagePresenter(router, interactor))
const homePage = NewHomeViewController(NewHomePagePresenter(router, interactor))
router.Initialize(homePage, loginPage)

router.GoToLoginPage()

//Inner layers must know nothing of outer layers' implementations
//Outer layers can depend on already-defined inner layers
//Inner layers can communicate to outer layers through interface

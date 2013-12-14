define [
  "devfeed",
  "login_view",
  "login_controller",
  "login_controller",
  "auth_app"
], (Devfeed, LoginView, LoginController, LogoutController, AuthApp) ->

  before ->
    Backbone.history.start
      pushState: true

  afterEach ->
    Devfeed.contentRegion.close()

  after ->
    Backbone.history.stop()

  describe "AuthApp", ->
    it "request auth:login:show displays login form and url", ->
      expect($('#content-region #login')[0]).to.not.exist
      Devfeed.trigger("auth:login:show")
      expect($('#content-region #login')[0]).to.exist
      expect(window.location.pathname).to.equal("/login")

  describe "AuthApp.Login.View", ->
    it "displays errors on empty form", ->
      loginView = new LoginView.LoginForm()
      loginView.on "auth:login", (data) ->
        loginView.triggerMethod("form:data:invalid", data)
      Devfeed.contentRegion.show(loginView)
      $("#login .signin").click()
      expect($("#login span.error").length).to.equal(2)

    it "displays alert on invalid credentials", ->
      loginView = new LoginView.LoginForm()
      loginView.on "auth:login", (data) ->
        loginView.triggerMethod("form:data:invalid", data)
      Devfeed.contentRegion.show(loginView)
      loginView.ui.emailInput.val("foo@bar.com")
      loginView.ui.passwordInput.val("secret")
      $("#login .signin").click()
      expect($("#login-alert-region .alert-box")).to.exist

  describe "AuthApp.Login.Controller", ->
    it "showLogin displays login form", ->
      expect($('#content-region #login')[0]).to.not.exist
      LoginController.showLogin()
      expect($('#content-region #login')[0]).to.exist

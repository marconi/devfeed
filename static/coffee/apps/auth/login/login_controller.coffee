define ["devfeed", "login_view", "websocket"], (Devfeed, LoginView, WebSocket) ->

  Devfeed.module "AuthApp.Login", (Login, Devfeed, Backbone, Marionette, $, _) ->

    Login.Controller =
      showLogin: ->
        loginView = new LoginView.LoginForm()
        loginView.on "auth:login", (data) ->
          loggingIn = Devfeed.request("session:login", data.email, data.password)
          $.when(loggingIn).done (result) ->
            if result and result.errors?
              loginView.triggerMethod("form:data:invalid", result)
            else
              Devfeed.execute("websocket:create")
              Devfeed.execute("show:userinfo")
              Devfeed.trigger("projects:list")
        Devfeed.contentRegion.show(loginView)

  return Devfeed.AuthApp.Login.Controller

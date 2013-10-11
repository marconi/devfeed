define ["devfeed", "header_controller"], (Devfeed, HeaderController) ->

  Devfeed.module "HeaderApp", (HeaderApp, Devfeed, Backbone, Marionette, $, _) ->

    API =
      showHeader: ->
        HeaderController.showHeader()

      removeUserinfo: ->
        HeaderController.removeUserinfo()

      showUserinfo: ->
        HeaderController.showUserinfo()

    Devfeed.on "auth:logout", ->
      API.removeUserinfo()

    Devfeed.commands.setHandler "show:userinfo", ->
      API.showUserinfo()

    HeaderApp.on "start", ->
      API.showHeader()
      loggingIn = Devfeed.request("session:isloggedin")
      $.when(loggingIn).done (isLoggedIn) ->
        if isLoggedIn
          API.showUserinfo()

  return Devfeed.HeaderApp

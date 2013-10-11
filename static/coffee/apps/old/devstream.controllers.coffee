define ['app'], (App) ->

  Controllers = {}

  #############################################################################
  ## Common controller
  #############################################################################

  Controllers.Common =
    showLogin: ->
      # loginView = new AppViews.LoginView
      # App.contentRegion.show(loginView)
      console.log "showing login form..."

  return Controllers

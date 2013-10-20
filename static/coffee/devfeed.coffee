define ["backbone", "marionette", "underscore"], (Backbone, Marionette, _) ->

  Devfeed = new Marionette.Application()

  Devfeed.addRegions
    headerRegion: "#header-region"
    contentRegion: "#content-region"

  Devfeed.navigate = (route, options) ->
    Backbone.history.navigate(route, options or {})

  Devfeed.redirect = (route) ->
    window.location = route

  Devfeed.getCurrentRoute = ->
    return Backbone.history.fragment

  Devfeed.messages =
    500: "Something went wrong while trying to process your request and we are already looking at it"

  Devfeed.isLoggedIn = (authCallback=null, unauthCallback=null) ->
    loggingIn = Devfeed.request("session:isloggedin")
    $.when(loggingIn).done (isLoggedIn) ->
      if isLoggedIn and authCallback
        authCallback()
      else if not isLoggedIn and unauthCallback 
        unauthCallback()

  Devfeed.on "home", ->
    Devfeed.isLoggedIn(
      -> Devfeed.trigger("projects:list"),
      -> Devfeed.trigger("auth:login:show"))

  Devfeed.on "initialize:after", ->
    $(document).foundation()
    $('#notification').miniNotification
      time: 5000
      opacity: 1
      closeButton: true
      closeButtonText: "&times;"

    if Backbone.history
      Backbone.history.start pushState: true

    if Devfeed.getCurrentRoute() == ""
      Devfeed.trigger("home")

  return Devfeed

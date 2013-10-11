define [
  'jquery',
  'backbone',
  'app',
  'app_models',
  'app_routers',
  'project_controller'
], (
  $,
  Backbone,
  App,
  AppModels,
  AppRouters,
  ProjectController
) ->

  #############################################################################
  ## App initialization
  #############################################################################

  App.on 'initialize:before', ->
    # set the router
    App.commonRouter = new AppRouters.Common()

    # create user session
    App.session = new AppModels.Session

  App.on 'initialize:after', ->
    Backbone.history.start
      pushState: true

    $(document).foundation()

  App.on 'start', ->
    # if not authenticated, navigate to login form
    if not App.session.isLoggedIn()
      App.commonRouter.controller.showLogin()

    # ProjectController.chatSubscribe()

  return App

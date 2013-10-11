define ['marionette', 'app_controllers'], (Marionette, AppControllers) ->

  Routers = {}

  #############################################################################
  ## Common router
  #############################################################################

  class Routers.Common extends Marionette.AppRouter
    controller: AppControllers.Common
    appRoutes:
      'login': 'showLogin'

  return Routers

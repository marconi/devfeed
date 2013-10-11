define ['backbone', 'app'], (Backbone, App) ->

  Models = {}

  #############################################################################
  ## UserSession
  #############################################################################

  class Models.Session extends Backbone.Model
    defaults:
      sessionId: null

    isAuthenticated: ->
      # if we don't have session id, fetch from backend
      if not @get('sessionId')
        $.ajax
          async: false
          type: 'GET'
          url: "/authenticated"
          success: (sessionId, textStatus, xhr) =>
            @set("sessionId", sessionId)

      return Boolean(@get('sessionId'))

  return Models

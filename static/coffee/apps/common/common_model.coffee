define ["devfeed"], (Devfeed) ->

  Devfeed.module "Common.Model", (Model, Devfeed, Backbone, Marionette, $, _) ->

    class Model.BaseModel extends Backbone.Model
      parse: (response, options) ->
        if response.s? and response.d? and response.s == 200
          return response.d
        return response

    class Model.BaseCollection extends Backbone.Collection
      comparator: (model) ->
        return model.get("id")
      parse: (response, options) ->
        if response.s? and response.d? and response.s == 200
          return response.d
        return response

  return Devfeed.Common.Model

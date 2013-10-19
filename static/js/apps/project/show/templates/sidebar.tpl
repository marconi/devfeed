<div id="project-name" class="row collapse">
  Footy Sprofila <a href="#" id="hide-sidebar">&laquo;</a>
</div>
<div id="find-task" class="row collapse">
  <div class="small-12 columns">
    <form>
      <div class="settings-cog"><a href="#"><i class="icon-cog"></i></a></div>
      <div class="keyword"><input type="text" placeholder="Find story..."></div>
    </form>
  </div>
</div>
<div id="tasks" class="row collapse">
  <div class="small-12 columns">
    <div class="inner">
      <ul>
        <% if (stories.length > 0) { %>
          <% stories.each(function(story) { %>
            <li>
              <a href="/projects/<%= id %>/stories/<%= story.get("id") %>">
                <%= story.get("name") %>
              </a>
            </li>
          <% }); %>
        <% } %>
      </ul>
    </div>
  </div>
</div>
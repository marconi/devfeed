<% if(!issynced) { %>
  <a href="#" class="disabled">
    Footy Sprofila <i class="syncing"></i>
  </a>
<% } else { %>
  <a href="/projects/<%= id %>">
    Footy Sprofila <i class="icon-play-sign"></i>
  </a>
<% } %>

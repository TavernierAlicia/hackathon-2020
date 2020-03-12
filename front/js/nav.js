function createNav() {
  document.write(`
    <!DOCTYPE html>
    <html lang="fr">
      <head>
        <meta charset="utf-8">
        <title>Recycle&Vous</title>
        <link rel="icon" href="/front/pictures/misc/favicon.png">
        <link rel="stylesheet" href="/front/css/main.css">
        <link rel="stylesheet" href="/front/css/bootstrap.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
        <script src="https://code.jquery.com/jquery-3.4.1.slim.min.js" integrity="sha384-J6qa4849blE2+poT4WnyKhv5vZF5SrPo0iEjwBvKU7imGFAV0wwj1yYfoRSJoZ+n" crossorigin="anonymous"></script>
        <script src="https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js" integrity="sha384-Q6E9RHvbIyZFJoft+2mJbHaEWldlvI9IOYy5n3zV9zzTtmI3UksdQRVvoxMfooAo" crossorigin="anonymous"></script>
        <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/js/bootstrap.min.js" integrity="sha384-wfSDF2E50Y2D1uUdj0O3uMBJnjuUD4Ih7YwaYd1iqfktj0Uod8GCExl3Og8ifwB6" crossorigin="anonymous"></script>
      </head>
      <body>
        <nav class="navbar sticky-top navbar-expand-lg navbar transparent">
          <a class="navbar-brand" href="/index"><img class="icon" src="/front/pictures/misc/icon.png"></a>
          <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
          </button>
          <div class="collapse navbar-collapse" id="navbarNav">
            <ul class="navbar-nav">
              <li class="nav-item active">
                <a class="nav-link" href="/recycleetvous">Mon Objet</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="/give">Je donne !</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="/workshop">Ateliers</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="/shop">Boutique</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="/reinsertion">Reinsertion</a>
              </li>
              <li class="nav-item">
              <a class="nav-link" href="/about">À propos</a>
            </li>
            </ul>
          </div>
        </nav>
      </body>
    </html>
  `);
}

createNav();


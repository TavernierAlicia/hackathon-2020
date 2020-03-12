function createFooter() {
  document.write(`
  <!DOCTYPE html>
  <html lang="fr">
    <head>
      <meta charset="utf-8">
      <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
    </head>
    <body>
    </body>
    <footer>
      <div class="container-fluid">
        <div class="container">
          <div class="row">
            <div class="col-4 footerlink">
            <h2 class="footer">RECYCLE & VOUS</h2>
              <a href="/about" >Le concept</a></br>
              <a href="/give" >Je donne!</a></br>
              <a href="/workshops" >Les ateliers</a></br>
              <a href="/shop" >La boutique</a></br>
              <a href="/reinsertion" >Reinserrez-vous!</a></br>

            </div>
            <div class="col-6 footerlink">
              <a href="#"><img src="/front/pictures/misc/Places.png">111 Rue de Paris, 75000 Paris</a></br>
              <a href="#" ><img src="/front/pictures/misc/phone.png">+330145678903</a></br>
              <a href="#" ><img src="/front/pictures/misc/messages.png">+recycleetvous.service@gmail.com</a></br>
              <a href="#" ><img src="/front/pictures/misc/newsletter.png">Abonnez-vous a la newsletter</a></br>
              <input type="text" placeholder="votre mail">
            </div>
            <div class="col-2 footerlink">
            <a href="#" ><img src="/front/pictures/misc/facebook.png">Facebook</a></br>
            <a href="#" ><img src="/front/pictures/misc/twitter.png">Twitter</a></br>
            <a href="#" ><img src="/front/pictures/misc/insta.png">Instagram</a></br>
            </div>
          </div>
          </div>
        <div class="row copyright">
          <p style="text-align:center">Ressource&Me - Copyright 2020</p></br>
          <p style="text-align:center"><a href="/legal"> Mentions légales </a></p>
        </div>
      </div>
    </footer>
  </html>
  `);

}

createFooter();
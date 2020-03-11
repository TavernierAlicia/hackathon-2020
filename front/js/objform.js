
function addDateNear() {
  var dateField = document.querySelector('div.dateNear');
  var today = new Date();
  var date = (today.getFullYear()+'-'+(today.getMonth()+1)+'-'+today.getDate()).toString()
  var datemax = (today.getFullYear()+'-'+(today.getMonth()+2)+'-'+today.getDate()).toString()
  dateField.innerHTML += `<h2>Je viens chercher l'objet le...   <input type="date" id="pickupdate" name="trip-start" value="`+date+`" min="`+date+`" max="`+datemax+`"></h2>`;
}



function addObjForm() {
  var numObject = document.querySelectorAll('div.object').length;

  if (numObject < 3) {
    console.log(numObject);
    document.getElementById('add').innerHTML += `<div class="object">
    <div class="form-group row">
      <div class="col-sm-12">
        <input type="text" name="objectname" class="form-control" id="objectname"
          placeholder="Nom de l'objet">
        </div>
      </div>

      <div class="form-group row">
        <label for="inlineFormCustomSelectPref" class="col-sm-2 col-form-label">Type d'objet:</label>
        <div class="col-sm-10">
          <select class="custom-select" name="subject" id="inlineFormCustomSelectPref">
            <option value="Textile">Textile (vêtements, linge de maison, chaussures..)</option>
            <option value="Vaisselle">Vaisselle (assiette, couverts, ustensiles..)</option>
            <option value="Jouets">Jouets (jeux de société, figurines, peluches..)</option>
            <option value="Meubles">Meubles (canapé,table, chaise, étagère...)</option>
            <option value="Électronique">Électronique (Ordinateur, Téléphone, Luminaires…) </option>
            <option value="Sports">Sports (Raquettes, Ballons, Tenues de sport..)</option>
            <option value="Décoration">Décoration (cadre, bibelot..)</option>
            <option value="Quincaillerie/Outillage">Quincaillerie/Ouitillage</option>
          </select>
        </div>
      </div>

      <div class="form-group row">
        <label for="inlineFormCustomSelect" class="col-sm-2 col-form-label">Etat de l'objet:</label>
        <div class="col-sm-10">
          <select class="custom-select" name="objectstate" id="inlineFormCustomSelect">
            <option value="Neuf">Neuf</option>
            <option value="Comme neuf">Comme neuf</option>
            <option value="Bon etat">Bon etat general</option>
            <option value="Use">Use</option>
            <option value="A reparer">A reparer</option>
          </select>
        </div>
      </div>

      <div class="form-group row">
        <label for="name" class="col-sm-2 col-form-label">Taille de l'objet (cm):</label>
        <div class="col-sm-10">
          <input type="text" name="objectsize" class="form-control" id="objectsize" placeholder="ex: 23">
        </div>
      </div>

      <div class="form-group row">
        <label for="name" class="col-sm-2 col-form-label">Poids de l'objet (g):</label>
        <div class="col-sm-10">
          <input type="text" name="objectweigth" class="form-control" id="objectweigth" placeholder="ex: 590">
        </div>
      </div>
  </div>
  <div class="beautyfy" </div> `;
  } else {
    alert("Vous pouvez apporter 3 objets maximum a cet evenement!");
  }
}
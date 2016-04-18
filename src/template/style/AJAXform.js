$(document).ready(function() {
    $('#form1').on('submit', function(e) {
        e.preventDefault(); // J'empêche le comportement par défaut du navigateur, c-à-d de soumettre le formulaire
 
        var $this = $(this); // L'objet jQuery du formulaire
 
        // récupération des données
        var date = $('#date').val();
        var nomcontact = $('#nomcontact').val();
        var adresse= $('#adresse').val();
        var nomtc = $('#nomtc').val();
        var adresse = $('#adresse').val();
        var tel = $('#tel').val();
        var auteurfiche = $('#auteurfiche').val();
        var lien = $('#lien').val();
        var circonstances = $('#circonstances').val();
        var priseenchargemedicale = $('#priseenchargemedicale').val();
        var priseenchargesociale = $('#priseenchargesociale').val();
        var demarches = $('#demarches').val();
        var problemes = $('#problemes').val();
        var attentes = $('#attentes').val();
        var observations = $('#observations').val();

        //Seuls le nom du contact et l'auteur sont obligatoires, on met donc une valeur
        //par défaut si les autres sont vide
        var default_val = 'Vide'
        if(date==='')
            date = default_val
        if(nomtc==='')
            nomtc = default_val
        if(adresse==='')
            adresse = default_val
        if(tel==='')
            tel = default_val
        if(lien==='')
            lien = default_val
        if(circonstances==='')
            circonstances = default_val
        if(priseenchargemedicale==='')
            priseenchargemedicale = default_val
        if(priseenchargesociale==='')
            priseenchargesociale = default_val
        if(demarches==='')
            demarches = default_val
        if(problemes==='')
            problemes = default_val
        if(attentes==='')
            attentes = default_val
        if(observations==='')
            observations = default_val
 
       //Test si champs obligatoires non remplis
        if(nomcontact === '' || auteurfiche === '') {
            alert('Il n\' y a pas de nom ni d\'auteur de la fiche, veuillez reneigner ces champs');
        } else {
            // Envoi de la requête HTTP en mode asynchrone
            $.ajax({
                url: $this.attr('/post'), // Le nom du fichier indiqué dans le formulaire
                type: $this.attr('post'), // La méthode indiquée dans le formulaire (get ou post)
                data: $this.serialize(), // Je sérialise les données (j'envoie toutes les valeurs présentes dans le formulaire)
                success: function(html) { // Je récupère la réponse du fichier PHP
                    alert(html); // J'affiche cette réponse
                }
            });
        }
    });
});
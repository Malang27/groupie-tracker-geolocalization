ymaps.ready(function () {
    var MaCarte = new ymaps.Map('map', {
        center: [51.15296161002877,71.48300799999997],
        zoom: 5,
        searchControlProvider: 'yandex#search'
    });
    var alem = new ymaps.Placemark([51.088352098944235,71.41420656258758]);
    MaCarte.geoObjects.add(alem);

    var maCollection = new ymaps.GeoObjectCollection();
    var npm = document.getElementsByClassName('placemark');
    myGeocoder = ymaps.geocode("astana");
    for (var i = 0; i < npm.length; ++i) {
        // Faites quelque chose avec npm[i], qui est un élément <li>
        myGeocoder = ymaps.geocode(npm[i].textContent.replace(/[-_]/g, " "));
        console.log(npm[i].textContent.replace(/[-_]/g, " "))
        myGeocoder.then(
            function (res) {
                MaCarte.geoObjects.add(res.geoObjects.get(1));
                // Afficher dans la console les données obtenues lors du géocodage de l'objet.
                console.log(res.geoObjects.get(1).properties.get('metaDataProperty').getAll());
            },
            function (err) {
                // Gérer l'erreur.
            }
        );
    }
    MaCarte.geoObjects.add(maCollection);
});

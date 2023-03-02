window.addEventListener("load", function load(event){

    var select_el = document.getElementById("placetypes");
    var ancestors_el = document.getElementById("ancestors");
    var descendants_el = document.getElementById("descendants");    

    var set_descendants = function(pt){
	
	sfomuseum_placetypes_descendants(pt, "common,optional,common_optional")
	    .then((data) => {
		
		var placetypes = JSON.parse(data);
		var count = placetypes.length;
		
		for (var i=0; i < count; i++) {
		    
		    var pt = placetypes[i];
		    var name = pt["name"];
		    
		    var item = document.createElement("li");
		    item.appendChild(document.createTextNode(name));
		    descendants_el.appendChild(item);
		}
		
	    }).catch((err) => {
		console.log("SAD", err);
	    });
	
    };

    var set_ancestors = function(pt){
	
	sfomuseum_placetypes_ancestors(pt, "common,optional,common_optional")
	    .then((data) => {
		
		var placetypes = JSON.parse(data);
		var count = placetypes.length;
		
		for (var i=0; i < count; i++) {
		    
		    var pt = placetypes[i];
		    var name = pt["name"];
		    
		    var item = document.createElement("li");
		    item.appendChild(document.createTextNode(name));
		    ancestors_el.appendChild(item);
		}
		
	    }).catch((err) => {
		console.log("SAD", err);
	    });
    };

    sfomuseum.placetypes.wasm.init().then(rsp => {

	sfomuseum_placetypes()
	    .then((data) => {

		var names = [];
		
		var placetypes = JSON.parse(data);
		var count = placetypes.length;
		
		for (var i=0; i < count; i++){
		    var pt = placetypes[i];
		    var name = pt["name"];
		    names.push(name);
		}

		names.sort()

		for (var i=0; i < count; i++){

		    var name = names[i];
		    
		    var opt = document.createElement("option");
		    opt.setAttribute("value", name);
		    opt.appendChild(document.createTextNode(name));

		    select_el.appendChild(opt);
		}

		/*

		   The underlying Go functions for these don't work as expected in a SFO Museum context
		   
		select_el.onchange = function(){

		    ancestors_el.innerHTML = "";
		    descendants_el.innerHTML = "";

		    var pt = select_el.value;

		    if (pt == ""){
			return;
		    }

		    set_descendants(pt);
		    set_ancestors(pt);		    
		};

		*/
	    })
	    .catch ((err) => {
		console.log("SAD", err);
	    });
	
    });
    
});

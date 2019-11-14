let nodes;
let edges;
let network;
let options = {
    layout: {
        hierarchical: false
    },
    physics: {
        maxVelocity: 15,
        barnesHut: {
            springLength: 500
        }
    },
    interaction: {hover: true},
    manipulation: {
        enabled: true
    }
};

let systemModel;

function reload() {
    let token = document.getElementById("github_token").value;
    $.getJSON('/getDiagramData?GITHUB_TOKEN=' + token, function (dataSet) {
        if (dataSet === null) {
            console.warn("Dataset came empty.");
            return;
        }
        nodes = new vis.DataSet(dataSet.nodes);
        // create an array with edges
        edges = new vis.DataSet(dataSet.edges);
    })
        .done(function () {
            // create a network
            let data = {
                nodes: nodes,
                edges: edges
            };
            const container = document.getElementById("mynetwork");
            network = new vis.Network(container, data, options);
        })
        .fail(function (response, textStatus, _) {
            console.log("Request failed: [" + textStatus + "] " + response.responseJSON.message)
        });
}

function removeSystemModel() {
    systemModel = nodes.get("system-model");
    nodes.delete("system-model");
}

function addSystemModel() {
    nodes.add(systemModel)
}

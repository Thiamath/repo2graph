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
        },
        stabilization: {
            fit: true
        }
    },
    interaction: {hover: true},
    manipulation: {
        enabled: false
    }
};

function reload() {
    let token = document.getElementById("github_token").value;
    $.getJSON('/getDiagramData?GITHUB_TOKEN=' + token, function (dataSet) {
        if (dataSet === null) {
            console.warn("Dataset came empty.");
            return;
        }
        nodelist.nodes = dataSet.nodes;
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

let deletedNodes = new Map();

function toggleNode(input) {
    if (input.checked) {
        nodes.add(deletedNodes.get(input.value));
        deletedNodes.delete(input.value);
        network.fit();
    } else {
        deletedNodes.set(input.value, nodes.get(input.value));
        let deleted = nodes.remove(input.value);
    }
}

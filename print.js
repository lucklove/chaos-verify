async function print(url) {
    const resp = await fetch(url);
    const json = await resp.json();
    const metric = json.data.result[1].values.map(x => parseFloat(x[1]));
    output = ""
    for(let i = 3; i < metric.length; i++) {
        const diff = (metric[i] - metric[i-1]) / metric[i-1];
        if(diff > 0.1 || diff < -0.1) {
            output += `${metric[i-3]},${metric[i-2]},${metric[i-1]},${metric[i]},1` + '\n';
        } else {
            output += `${metric[i-3]},${metric[i-2]},${metric[i-1]},${metric[i]},0` + '\n';
        }
    }
    console.log(output);
}

async function metric(url) {
    const resp = await fetch(url);
    const json = await resp.json();
    const metric = json.data.result[1].values.map(x => parseFloat(x[1]));
    console.log(metric.join(","));
}
console.log("background");
function handleActivated(tab) {
    browser.tabs.get(tab.tabId, e=>{
        console.log(e.url)
    })
    browser.tabs.executeScript({
        code:`var x=document.getElementsByTagName("h1");
for(let i=0; i<x.length;i++){
    console.log(5);
}`,
    })
}


browser.tabs.onActivated.addListener(handleActivated);

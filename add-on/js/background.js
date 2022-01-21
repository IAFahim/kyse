console.log("background");
function handleActivated(tab) {
    browser.tabs.get(tab.tabId, e=>{
        console.log(e.url)
    })
    browser.tabs.executeScript({
        file:"js/foreground.js"
    });
}


browser.tabs.onActivated.addListener(handleActivated);

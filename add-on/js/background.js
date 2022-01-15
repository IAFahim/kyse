console.log("background");
function handleActivated(tab) {
    browser.tabs.get(tab.tabId, e=>{
        console.log(e.url)
    })
}

browser.tabs.onActivated.addListener(handleActivated);

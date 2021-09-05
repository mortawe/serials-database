import {config} from "./pages/config.js";
import {menuPage} from "./pages/Menu/Menu.js";
export const application = document.getElementById('app');

application.addEventListener('click', e => {
    const {target} = e;

    if (target instanceof HTMLAnchorElement) {
        e.preventDefault();
        config[target.dataset.section].open(target.href);
    }
});

menuPage();
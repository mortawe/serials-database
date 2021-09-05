import {application} from "../../main.js";
import {createBack, createRef} from "../../utils/back.js";

export function listGenrePage(href, data) {
    const id = parseInt(/genre\/(\d+)/.exec(href)[1]);
    if (!id) {
        console.error("no id");
        return;
    }
    application.innerHTML = '';
    const section = document.createElement('section');
    section.dataset.sectionName = "genre/list";

    const header = document.createElement('h1');
    header.textContent = "Genre List";
    section.appendChild(header);

    const back = createBack();
    section.appendChild(back);
    if (!data) {
        HttpModule.post({
            url: '/genre/get',
            body: {id: id},
            callback: (status, response) => {
                switch (status) {
                    case 200: {
                        const data = JSON.parse(response)
                        data.forEach((elem) => {
                            const a = createRef(elem.title, `/show/${elem.show_id}`, 'show/update')
                            section.appendChild(a);
                            return elem;
                        });
                        break;
                    }
                    default: {
                        const error = JSON.parse(response);
                        alert(error);
                    }
                }
            }
        });
    }
    application.appendChild(section)
}
import {createRef} from "../../utils/back.js";
import {TableComponent} from "../../components/Table/Table.js";
import {application} from "../../main.js";
import {menuPage} from "../Menu/Menu.js";

export function tablePage(type, data) {
    application.innerText = '';
    menuPage();

    const section = document.createElement('section');
    section.dataset.sectionName = type.name;

    const header = document.createElement('h1');
    header.textContent = type.header;
    section.appendChild(header);
    const ref = createRef(`Create ${type.header}`, `${type.name}/create`, `${type.name}/create`)
    section.appendChild(ref);
    const refSearch = createRef(`Search ${type.header}`, `${type.name}/find`, `${type.name}/find`)
    section.appendChild(refSearch);

    const tableNode = document.createElement('table');

    if (data) {
        data.forEach(elem => {
            elem.release = new Date(elem.release).toDateString();
            elem.birthdate = new Date(elem.birthdate).toDateString();
        })
        const table = new TableComponent({
            tmplName: type.tmplName,
            parent: tableNode,
            data: data,
        });
        table.render();
        section.appendChild(tableNode);
    } else {
        const em = document.createElement('em');
        em.textContent = 'Loading...';
        section.appendChild(em);

        HttpModule.post(type.request);
    }
    application.appendChild(section);
}
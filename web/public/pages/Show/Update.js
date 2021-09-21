import {application} from "../../main.js";
import {FormComponent} from "../../components/Form/Form.js";
import {TABLE_MAP} from "../Table/map.js";
import {tablePage} from "../Table/Table.js";
import {menuPage} from "../Menu/Menu.js";

export function updateShowPage(href, data) {
    application.innerText = '';
    menuPage();

    const id = parseInt(/show\/(\d+)\//.exec(href)[1]);
    if (!id) {
        console.error("no id");
        return;
    }
    if (!data) {
        HttpModule.post({
            url: "/show/get",
            body: {"id": id},
            callback: (status, response) => {
                switch (status) {
                    case 200: {
                        const data = JSON.parse(response);
                        data.release = new Date(data.release).toISOString().substr(0, 10);
                        HttpModule.post({
                            url: "/person/getAll",
                            callback: (status, response) => {
                                switch (status) {
                                    case 200:
                                        data.personList = JSON.parse(response);
                                        updateShowPage(href, data);
                                        break;
                                    default:
                                        const error = JSON.parse(response);
                                        alert(error);
                                }
                            }
                        })
                        break;
                    }
                    default:
                        const error = JSON.parse(response);
                        alert(error);
                        return;
                }
            }
        })
    }
    const section = document.createElement('section');
    section.dataset.sectionName = "updateShow";

    const header = document.createElement('h1');
    header.textContent = "Update Show";
    section.appendChild(header);


    if (data) {
        const selectedIDs = data.person.map(elem => elem.person_id)
        console.log('selected' + selectedIDs)
        console.log(data.personList)
        data.personList.map(elem => {
                if (selectedIDs.includes(elem.person_id)) {
                    elem.selected = "selected";
                } else {
                    elem.selected = "";
                }
                return elem;
            }
        )
    }
    const formNode = document.createElement('form');
    const form = new FormComponent({
        tmplName: "Show/Update.mustache",
        parent: formNode,
        config: data,
    });
    form.render();
    section.appendChild(formNode);


    formNode.addEventListener('submit', (evt) => {
        evt.preventDefault();

        const title = document.getElementById('title').value;
        const release = document.getElementById('release').value;
        const description = document.getElementById('description').value;
        const episode_num = parseInt(document.getElementById('episode_num').value);
        const genre = document.getElementById('genre').value;

        function getSelectValues(select) {
            const result = [];
            const options = select && select.options;
            let opt;

            let i = 0, iLen = options.length;
            for (; i < iLen; i++) {
                opt = options[i];

                if (opt.selected) {
                    result.push(opt.value || opt.text);
                }
            }
            return result;
        }

        const persons = getSelectValues(document.getElementById('selector')).map(function (elem) {
            return {'person_id': parseInt(elem)}
        });

        HttpModule.post({
            url: '/show/update',
            body: {
                show_id: id,
                title: title,
                release: new Date(release),
                description: description,
                person: persons,
                genre: genre,
                episode_num: episode_num
            },
            callback: (status, response) => {
                switch (status) {
                    case 200: {
                        tablePage(TABLE_MAP.show);
                        break;
                    }
                    default:
                        const error = JSON.parse(response);
                        alert(error);
                }
            }
        })
    });
    application.appendChild(section);
}
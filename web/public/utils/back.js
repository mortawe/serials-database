export function createBack() {
    const back = document.createElement('a');
    back.href = '/';
    back.textContent = 'Menu';
    back.dataset.section = 'menu';

    return back;
}

export function createRef(name, path, section) {
    const ref = document.createElement('a');
    ref.href = path;
    ref.textContent = name;
    ref.dataset.section = section;

    return ref;
}
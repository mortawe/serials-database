
export function createRef(name, path, section) {
    const ref = document.createElement('a');
    ref.href = path;
    ref.textContent = name;
    ref.dataset.section = section;

    return ref;
}
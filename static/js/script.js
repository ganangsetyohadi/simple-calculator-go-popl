document.getElementById('calcForm').addEventListener('submit', function (e) {
    e.preventDefault();

    const form = e.target;
    const formData = new FormData(form);

    fetch('/', {
        method: 'POST',
        body: formData
    })
    .then(response => response.text())
    .then(html => {
        const parser = new DOMParser();
        const doc = parser.parseFromString(html, 'text/html');
        const result = doc.getElementById('result').textContent;
        document.getElementById('result').textContent = result;
    });
});

function toggleSecondInput() {
    const operator = document.getElementById("operator").value;
    const secondNumberGroup = document.getElementById("secondNumberGroup");

    if (operator === "sqrt") {
        secondNumberGroup.style.display = "none";
    } else {
        secondNumberGroup.style.display = "block";
    }
}

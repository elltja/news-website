const resizer = document.getElementById("main-resizer");
const leftPane = document.querySelector(".write-post");
const rightPane = document.querySelector(".lists");
const container = document.getElementById("main-wrapper");

let isDragging = false;

resizer.addEventListener("mousedown", (e) => {
  e.preventDefault();
  isDragging = true;
  document.body.style.cursor = "col-resize";
});

document.addEventListener("mousemove", (e) => {
  if (!isDragging) return;
  const containerLeft = container.getBoundingClientRect().left;
  const containerWidth = container.offsetWidth;
  const offsetX = e.clientX - containerLeft;

  const leftPercent = (offsetX / containerWidth) * 100;
  const rightPercent = 100 - leftPercent;

  leftPane.style.width = `${leftPercent}%`;
  rightPane.style.width = `${rightPercent}%`;
});

document.addEventListener("mouseup", () => {
  isDragging = false;
  document.body.style.cursor = "default";
});

const form = document.querySelector("section.write-post form");

form.addEventListener("submit", async (e) => {
  e.preventDefault();
  const titleElement = document.querySelector("input[name='title']");
  const contentElement = document.querySelector("textarea[name='content']");
  const title = titleElement.value;
  const content = contentElement.value;
  if (title == "" || content == "") {
    displayError("All fields are required");
    return;
  }

  const res = await fetch("/api/admin/create-article", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
    body: JSON.stringify({ title, content }),
  });
  if (!res.ok) {
    displayError(await res.text());
    return;
  }
  titleElement.value = "";
  contentElement.value = "";
  displaySucces(await res.text());
});

function displayError(msg) {
  document.getElementById("error-msg").innerText = msg;
}

function displaySucces(msg) {
  document.getElementById("success-msg").innerText = msg;
}

document.querySelectorAll(".delete-article-btn").forEach((btn) => {
  btn.addEventListener("click", async () => {
    const articleId = btn.getAttribute("data-article-id");
    await fetch(`/api/admin/delete-article/${articleId}`, {
      method: "DELETE",
      credentials: "include",
    });
    window.location.reload();
  });
});

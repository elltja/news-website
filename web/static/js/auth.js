const form = document.getElementById("auth-form");

form.addEventListener("submit", async (e) => {
  e.preventDefault();

  const email = document.querySelector("input[name='email']").value;
  const password = document.querySelector("input[name='password']").value;
  console.log({ email, password });

  const res = await fetch("/api/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ email, password }),
  });
  if (!res.ok) {
    displayError(await res.text());
  }
  window.location.pathname = "/";
});

function displayError(msg) {
  document.getElementById("error-msg").innerText = msg;
}

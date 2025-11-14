document.addEventListener("DOMContentLoaded", () => {
  const btnInvest = document.getElementById("invest");
  const containerProduct = document.getElementById("container-product");

  const confirmModal = document.getElementById("confirmModal");

  btnInvest.addEventListener("click", async () => {
    confirmModal.classList.toggle("hidden");
    const isApply = await isApplyPopup(confirmModal);

    if (isApply) {
      handleInfetch(containerProduct);
    }
  });
});
/**
 * tampilkan modal
 * return : boolean
 */

function isApplyPopup(confirmModal) {
  return new Promise((resolve) => {
    // agar tidak menghasilkan undifined

    const batal = confirmModal.querySelector("#action #cancelBtn");
    const ya = confirmModal.querySelector("#action #confirmBtn");

    batal.onclick = () => {
      confirmModal.classList.add("hidden");
      resolve(false);
    };

    ya.onclick = () => {
      confirmModal.classList.add("hidden");
      resolve(true);
    };
  });
}

/*
 * saat tombol invest di click , maka data di kirim ke server
 * */
function handleInfetch(containerProduct) {
  const productName = containerProduct.querySelector("h2").value;
  const total_product = containerProduct.querySelector(
    "input#total_product",
  ).value;

  const url = "/product/product-user";
  fetch(url, {
    headers: {
      "Content-Type": "application/json",
    },
    method: "POST",
    body: JSON.stringify({
      total_product,
      name: productName,
    }),
  });
}

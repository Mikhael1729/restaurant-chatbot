<template>
  <div class="title-description">
    <img src="../assets/logo.png"/>
    <h1>Tecnorestaurant</h1>

    <p>
      <i>
        Puedes escribirnos cualquier duda acerca de nuestros servicios por medio
        de este chat
      </i>
    </p>
  </div>
</template>

<style lang="scss" scoped>
.title-description {
  display: flex;
  flex-direction: column;
  height: fit-content;
  align-self: center;
  border: 1px solid var(--text-color);
  border-radius: 0.8em;
  text-align: center;
  align-items: center;

  img {
    width: 80px;
  }

  h1 {
    color: #dbdbdb;
    font-size: 2.5em;
  }

  p {
    margin-top: 0.5em;
    font-size: 1.2em;
    line-height: 170%;
  }
}
</style>

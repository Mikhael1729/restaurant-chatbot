<template>
  <div class="title-description">
    <h1 class="title">Restaurante<br />Chat</h1>

    <p class="description">
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
  border-radius: 0.5em;
  padding: 2em;
  color: red;
  width: 30%;

  .title {
    color: #dbdbdb;
  }

  .description {
    flex: 1;
    margin-top: 1em;
    font-size: 0.8em;
    line-height: 170%;
  }

  .logo {
    width: 80px;
    align-self: center;
  }
}
</style>

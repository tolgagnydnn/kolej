<template>
    <div class="wrapper">
        <nav class="navbar navbar-expand-lg navbar-light">
            <a class="navbar-brand" href="index.html">
                <img src="@/assets/images/logo.png" class="img-fluid deskLogo">
            </a>
            <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>

            <div class="collapse navbar-collapse" id="navbarSupportedContent">
                <ul class="navbar-nav ml-lg-auto">
                    <li class="nav-item">
                        <a href="#" class="nav-link">
                        <router-link to="/" name=”anasayfa” exact> Anasayfa </router-link>
                        </a>
                     </li>
                  </ul>
            </div>
          </nav>
      </div>
</template>

<script>

export default {
  data(){
    return {
      menuLinks : [
        {link: "Anasayfa"},
        {link: "Hakkımızda"},
        {link: "Bize Ulaşın"}
      ]
    }
  }

}

</script>

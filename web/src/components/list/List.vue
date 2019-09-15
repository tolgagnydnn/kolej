<template>
  <div class="container listContainer py-5">
    <h1> Okul Listesi </h1>
    <div class="row py-4">
      <!--ListBox Start -->
        <div class="col-lg-8 col-md-8 col-sm-8 col-12">
          <div class="listBox" v-for="list in schoolList">
            <div class="row">

              <div class="col-lg-12 col-md-12 col-sm-12 col-12">
                <div class="row">
                  <div class="col-lg-2 col-md-2 col-sm-2 col-12 d-flex align-items-center justify-content-center">
                    <div class="images">
                        <img src="@/assets/images/okul.png" class="img-fluid">
                    </div>
                  </div>
                  <!-- Content Start -->
                  <div class="col-lg-10 col-md-10 col-sm-10 col-12">
                      <div class="listContent">
                        <div class="row">

                          <div class="col-lg-9 col-md-9 col-sm-9 col-12">
                            <div class="schollName">
                              <a href=""> {{list.adi}} </a>
                            </div>
                          </div>

                          <div class="col-lg-3 col-md-3 col-sm-3 col-12">
                              <i class="fas fa-star"></i>
                              <i class="fas fa-star"></i>
                              <i class="fas fa-star"></i>
                          </div>

                          <div class="col-lg-12 col-md-12 col-sm-12">
                            <div class="row">
                              <div class="col-lg-8 col-md-8 col-sm-8 col-12">
                                <div class="schoolAdress">
                                  <p> {{list.adres}}  </p>
                                </div>
                              </div>
                              <div class="col-lg-4 col-md-4 col-sm-4 col-12">
                                <div class="schollWeb">
                                  <p href="#"> {{list.website}} </p>
                                </div>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  <!-- Content End -->
                  </div>
                </div>
              </div>

            </div>
          </div>
          <!-- Listbox End -->
        </div>

        <div class="col-lg-4 col-md-4 col-sm-4 col-12">
            <!-- Buraya Başka Compoentler gelecek -->
        </div>
    </div>
  </div>
</template>

<script>

import axios from 'axios';
export default {
  data(){
    return {
      schoolList : []
    }
  },
  mounted: function(){
      this.etkinlikListe()
  },
  methods: {
      etkinlikListe: function() {
          axios.get('http://localhost:8090/okul/liste')
          .then((result) => {
              this.schoolList = result.data
          });
      }
  }
}

</script>

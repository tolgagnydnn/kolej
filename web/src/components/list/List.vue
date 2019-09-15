<template>
  <div>
    <div class="card" v-for="list in schoolList" :key="list.id">
        <p> {{list.adi}} </p>
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

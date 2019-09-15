<template>
  <div class="container listContainer py-5">
    <h1> Okul Listesi </h1>
    <div class="row py-4">
          <!--ListBox Start -->
            <appList> </appList>
          <!-- Listbox End -->
    </div>

        <div class="col-lg-4 col-md-4 col-sm-4 col-12">
            <!-- Buraya Başka Compoentler gelecek -->
        </div>
    </div>
  </div>
</template>

<script>
import List from "./List";
export default {
  components: {
    "appList": List
  }
}

</script>

<template>
  <v-app>
    <v-navigation-drawer app :mini-variant="miniVariant">
      <v-list-item @click="miniVariant = !miniVariant" class="mini_swapper">
        <v-list-item-content>
          <v-list-item-title class="title">Das Kapital</v-list-item-title>
          <v-list-item-subtitle>subtext</v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>
      <v-divider></v-divider>
      <v-list dense nav>
        <v-list-item v-for="item in items" :key="item.title" link :to="item.link">
          <v-list-item-icon>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
      <template v-slot:append>
        <div class="pa-2">
          <v-btn block @click="logout"><v-icon>mdi-logout</v-icon>Logout</v-btn>
        </div>
      </template>
    </v-navigation-drawer>
    <v-app-bar id="app-bar" absolute app color="transparent" flat height="75">
      <v-app-bar-nav-icon @click="miniVariant = !miniVariant"></v-app-bar-nav-icon>
      <v-toolbar-title class="hidden-sm-and-down font-weight-light" v-text="$route.name"/>
      <v-spacer />
      <v-text-field :label="$t('search')" color="secondary" hide-details style="max-width: 165px;">
        <template v-if="$vuetify.breakpoint.mdAndUp" v-slot:append-outer>
          <v-btn class="mt-n2" elevation="1" fab small>
            <v-icon>mdi-magnify</v-icon>
          </v-btn>
        </template>
      </v-text-field>

      <div class="mx-3" />
      <v-btn class="ml-2" min-width="0" text to="/">
        <v-icon>mdi-view-dashboard</v-icon>
      </v-btn>

      <v-menu bottom left offset-y origin="top right" transition="scale-transition">
        <template v-slot:activator="{ attrs, on }">
          <v-btn class="ml-2" min-width="0" text v-bind="attrs" v-on="on">
            <v-badge color="red" overlap bordered>
              <template v-slot:badge>
                <span>5</span>
              </template>

              <v-icon>mdi-bell</v-icon>
            </v-badge>
          </v-btn>
        </template>

        <v-list :tile="false" nav>
          <div>
            <app-bar-item v-for="(n, i) in notifications" :key="`item-${i}`">
              <v-list-item-title v-text="n" />
            </app-bar-item>
          </div>
        </v-list>
      </v-menu>

      <v-btn class="ml-2" min-width="0" text to="/profile">
        <v-icon>mdi-account</v-icon>
      </v-btn>
    </v-app-bar>
    <v-main>
      <v-container fluid>
        <router-view></router-view>
        <v-footer id="dashboard-core-footer" app>
          <v-container>
            <v-row align="center" no-gutters>
              <v-col class="text-center mb-sm-0 mb-5" cols="auto"></v-col>
              <v-spacer class="hidden-sm-and-down" />

              <v-col cols="12" md="auto">
                <div class="body-1 font-weight-light pt-6 pt-md-0 text-center">
                  &copy; 2019, made with
                  <v-icon size="18">
                    mdi-heart
                  </v-icon>
                  by <a href="https://www.creative-tim.com">Creative Tim</a> for a better web.
                </div>
              </v-col>
            </v-row>
          </v-container>
        </v-footer>
      </v-container>
      <v-snackbar v-if="alertData.display"
                  :value="true"
                  bottom
                  :color="alertData.color"
                  outlined
                  right>{{ alertData.text }}</v-snackbar>
    </v-main>
  </v-app>
</template>

<script>
  export default {
    data () {
      return {
        bottomNav: null,
        miniVariant: false,
        drawer: true,
        value: true,
        items: [
          { title: 'Dashboard', icon: 'mdi-view-dashboard', link: '/' },
          { title: 'Statistic', icon: 'mdi-chart-areaspline', link: '/statistic' },
          { title: 'Profile', icon: 'mdi-account-box', link: '/profile'},
          { title: 'Categories', icon: 'mdi-format-list-bulleted-square', link: '/categories'},
        ],
        notifications: [
          'Mike John Responded to your email',
          'You have 5 new tasks',
          'You\'re now friends with Andrew',
          'Another Notification',
          'Another one',
        ],
      }
    },
    computed: {
      alertData() {
        return this.$store.state.alertData;
      }
    },
    created() {
      this.getUserCategories();
    },
    methods: {
      logout() {
        console.log('logout');
      },

      showAlertR: function(color, text, timeout) {
        this.alertText = text;
        this.alertColor = color;
        this.alertDisplay = true;
        setTimeout(() => {
          this.alertDisplay = false;
        }, +timeout * 1000);
      },

      getUserCategories() {
        if (this.$store.categories) {
          //return;
        }
        this.askBackend('user_category/get', {})
                .then(({data}) => {
                  if (data.ok) {
                    this.$store.commit('setCategories', data.categories || []);
                  }
                })
                .catch()
      }
    }
  }
</script>

<style scoped>
.mini_swapper{
  cursor: pointer;
}
</style>
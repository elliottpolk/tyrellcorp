<template>
  <b-card header-tag="header" bg-variant="light" class="mt-3 w-75">
    <template slot="header">
      <div class="d-flex flex-row font-weight-bold pt-0">Replicants</div>
    </template>

    <b-row></b-row>

    <b-row>
      <b-col>
        <b-table
          striped
          hover
          fixed
          bordered
          :items="displayed"
          :fields="fields"
          :sort-compare="sortCompare"
          :sort-by.sync="sortBy"
          :sort-desc.sync="sortDesc"
        >
          <template slot="name" slot-scope="row">
            <div :title="row.item.id">{{ row.item.name }}</div>
          </template>

          <template slot="show_details" slot-scope="row">
            <b-button variant="link" size="sm" @click="row.toggleDetails">
              <fontawesome :icon="row.detailsShowing ? 'caret-square-up' : 'caret-square-down'"/>
            </b-button>
          </template>

          <template slot="row-details" slot-scope="row">
            <b-card no-body class="p-4">
              <b-form-group
                label-cols-sm="3"
                label="ID:"
                label-align-sm="right"
                label-for="id-input"
              >
                <b-form-input id="id-input" type="text" v-model="row.item.id" class="w-75" disabled></b-form-input>
              </b-form-group>

              <b-form-group
                label-cols-sm="3"
                label="Repository:"
                label-align-sm="right"
                label-for="repository-input"
              >
                <b-form-input
                  id="repository-input"
                  type="text"
                  v-model="row.item.repository"
                  class="w-75"
                  disabled
                ></b-form-input>
              </b-form-group>

              <b-form-group label-cols-sm="3" label="Services:" label-align-sm="right" class="mb-0">
                <div class="d-flex justify-content-around pt-2">
                  <b-form-checkbox switch v-model="row.item.create" disabled>Create</b-form-checkbox>
                  <b-form-checkbox switch v-model="row.item.retrieve" disabled>Retrieve</b-form-checkbox>
                  <b-form-checkbox switch v-model="row.item.update" disabled>Update</b-form-checkbox>
                  <b-form-checkbox switch v-model="row.item.delete" disabled>Delete</b-form-checkbox>
                </div>
              </b-form-group>

              <b-form-group label-cols-sm="3" label="Fields:" label-align-sm="right" class="mb-0">
                <b-table
                  striped
                  hover
                  fixed
                  bordered
                  show-empty
                  :items="row.item.fields"
                  :fields="specFields"
                  class="tc-b-table"
                >
                  <template slot="empty">
                    <h5>No Available Fields</h5>
                  </template>

                  <template slot="show_details" slot-scope="row">
                    <b-button variant="link" size="sm" @click="row.toggleDetails">
                      <fontawesome
                        :icon="row.detailsShowing ? 'caret-square-up' : 'caret-square-down'"
                      />
                    </b-button>
                  </template>

                  <template slot="row-details" slot-scope="row">
                    <b-card no-body class="d-flex flex-row">
                      <b-col class="p-3">
                        <b-row class="px-3 mb-3 justify-content-between">
                          <div class="font-weight-bold">Description:</div>
                          <div class="text-muted">{{ row.item.description }}</div>
                        </b-row>

                        <b-row class="px-3 mb-3 justify-content-between">
                          <span class="font-weight-bold">Field Type:</span>
                          <span class="text-muted">{{ row.item.type }}</span>
                        </b-row>

                        <b-row class="px-3 mb-3 justify-content-between">
                          <span class="font-weight-bold">Sequence Number:</span>
                          <span class="text-muted">{{ row.item.sequence}}</span>
                        </b-row>

                        <b-row class="px-3 justify-content-between">
                          <b-col class="d-inline-flex flex-column p-0 align-items-start">
                            <b-form-checkbox switch disabled v-model="row.item.is_list">Is List?</b-form-checkbox>
                            <b-form-checkbox switch disabled v-model="row.item.is_key">Is Key?</b-form-checkbox>
                          </b-col>
                        </b-row>
                      </b-col>
                    </b-card>
                  </template>
                </b-table>
              </b-form-group>
            </b-card>
          </template>
        </b-table>
      </b-col>
    </b-row>
  </b-card>
</template>

<script>

export default {
  props: {
    parentState: Object
  },

  data () {
    return {
      fields: [
        { key: 'name', sortable: true },
        { key: 'package', sortable: true },
        { key: 'state', sortable: true },
        { key: 'show_details', label: 'Details', sortable: false }
      ],
      specFields: [
        { key: 'name', sortable: false },
        { key: 'type', sortable: false },
        { key: 'show_details', label: 'Details', sortable: false }
      ],
      replicants: [],
      displayed: [],
      sortBy: '',
      sortDesc: false,
      api: {
        baseUrl: '/api/v1/specs'
      }
    }
  },

  mounted () {

  },

  created () {
    this.$http.get(this.api.baseUrl, { request_id: this.$uuid.v4() }).then(res => {
      if (res.body.payload && res.body.payload !== null) {
        let replicants = []
        res.body.payload.forEach(el => {
          replicants.push(el)
        })

        this.parentState.replicants = replicants
      }
    }).catch(err => {
      console.error(err)
    })
  },

  watch: {
    'parentState.replicants': function (val) {
      console.debug('parent changed')
      this.displayed = this.replicants = this.parentState.replicants
    }
  },

  methods: {
    sortCompare (a, b, key) {
      // handle nested keys
      key.split('.').forEach(k => {
        a = a[k]
        b = b[k]
      })

      if (!a || a === null) {
        return 1
      }

      if (!b || b === null) {
        return 0
      }

      if (typeof a === 'number' && typeof b === 'number') {
        // if both compared fields are native numbers
        return a < b ? -1 : (a > b ? 1 : 0)
      } else {
        // stringify the field data and use String.localeCompare
        return a.localCompare(b, 'en', { sensitivity: 'base' })
      }
    }
  }
}
</script>

<style>
</style>

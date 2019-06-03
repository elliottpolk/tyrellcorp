<template>
  <b-card header-tag="header" footer-tag="footer" bg-variant="light" class="mt-3 w-75">
    <template slot="header">
      <div class="d-flex flex-row font-weight-bold pt-0">New Spec</div>
    </template>

    <b-form-group
      label-cols-sm="3"
      label="Username:"
      label-align-sm="right"
      label-for="username-input"
    >
      <b-form-input
        id="username-input"
        type="text"
        v-model="specForm.record_info.created_by"
        class="w-50"
        required
        :disabled="saving"
      ></b-form-input>
    </b-form-group>

    <b-form-group label-cols-sm="3" label="Name:" label-align-sm="right" label-for="name-input">
      <b-form-input
        id="name-input"
        type="text"
        v-model="specForm.name"
        class="w-50"
        required
        :disabled="saving"
      ></b-form-input>
    </b-form-group>

    <b-form-group
      label-cols-sm="3"
      label="Package:"
      label-align-sm="right"
      label-for="package-input"
    >
      <b-form-input
        id="package-input"
        type="text"
        v-model="specForm.package"
        class="w-50"
        required
        :disabled="saving"
      ></b-form-input>
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
        v-model="specForm.repository"
        class="w-50"
        required
        :disabled="saving"
      ></b-form-input>
    </b-form-group>

    <b-form-group label-cols-sm="3" label="Services:" label-align-sm="right" class="mb-0">
      <div class="d-flex justify-content-around pt-2">
        <b-form-checkbox switch v-model="specForm.create" :disabled="saving">Create</b-form-checkbox>
        <b-form-checkbox switch v-model="specForm.retrieve" :disabled="saving">Retrieve</b-form-checkbox>
        <b-form-checkbox switch v-model="specForm.update" :disabled="saving">Update</b-form-checkbox>
        <b-form-checkbox switch v-model="specForm.delete" :disabled="saving">Delete</b-form-checkbox>
      </div>
    </b-form-group>

    <b-form-group label-cols-sm="3" label="Fields:" label-align-sm="right" class="mb-0">
      <b-table
        striped
        hover
        fixed
        bordered
        show-empty
        :items="specForm.fields"
        :fields="fields"
        class="tc-b-table"
      >
        <template slot="empty">
          <h5>No Available Fields</h5>
        </template>

        <template slot="HEAD_show_details">
          <b-button
            variant="link"
            size="sm"
            style="line-height: .8rem;"
            @click="newField"
            :disabled="saving"
          >
            <fontawesome icon="plus-square"/>
          </b-button>
        </template>

        <template slot="show_details" slot-scope="row">
          <b-button variant="link" size="sm" @click="row.toggleDetails">
            <fontawesome :icon="row.detailsShowing ? 'caret-square-up' : 'caret-square-down'"/>
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

                <b-col class="d-inline-flex p-0 align-items-end flex-row-reverse">
                  <b-button size="sm" variant="link">
                    <fontawesome icon="trash"/>
                  </b-button>
                </b-col>
              </b-row>
            </b-col>
          </b-card>
        </template>
      </b-table>
    </b-form-group>

    <template slot="footer">
      <b-button-toolbar class="d-flex flex-row-reverse">
        <b-button-group size="md" class="mr-1">
          <b-button @click="resetSpecForm" :disabled="saving">Cancel</b-button>
          <b-button v-if="!saving" @click="saveSpec">Save</b-button>
          <b-button v-if="saving" disabled>
            <b-spinner small></b-spinner>
            <span class="m-1">Saving...</span>
          </b-button>
        </b-button-group>
      </b-button-toolbar>
    </template>

    <!-- ------------------------------------------------------------------ -->
    <!-- New Spec Field Modal                                               -->
    <!-- ------------------------------------------------------------------ -->
    <b-modal
      centered
      id="newFieldModal"
      @shown="resetFieldForm"
      @ok="handleFieldOk"
      title="New Field"
      hide-header-close
    >
      <b-form-group
        label-cols-sm="3"
        label="Name:"
        label-align-sm="right"
        label-for="field-name-input"
      >
        <b-form-input
          ref="field-name-input"
          id="field-name-input"
          type="text"
          v-model="fieldForm.name"
          class="w-75"
        ></b-form-input>
      </b-form-group>

      <b-form-group
        label-cols-sm="3"
        label="Description:"
        label-align-sm="right"
        label-for="field-description-input"
      >
        <b-form-textarea
          ref="field-description-input"
          id="field-description-input"
          v-model="fieldForm.description"
          class="w-75"
        ></b-form-textarea>
      </b-form-group>

      <b-form-group
        label-cols-sm="3"
        label="Field Type:"
        label-align-sm="right"
        label-for="field-type-selector"
      >
        <b-form-select
          ref="field-type-selector"
          id="field-type-selector"
          v-model="fieldForm.type"
          :options="fieldTypes"
        >
          <template slot="first">
            <option :value="null" selected>-- Select a type --</option>
          </template>
        </b-form-select>
      </b-form-group>

      <b-form-group label-cols-sm="3" label-align-sm="right" class="mb-0">
        <div class="d-flex justify-content-end pt-2 w-75">
          <b-form-checkbox switch v-model="fieldForm.is_list" class="ml-4">Is List</b-form-checkbox>
          <b-form-checkbox switch v-model="fieldForm.is_key" class="ml-4">Is Key</b-form-checkbox>
        </div>
      </b-form-group>
    </b-modal>
  </b-card>
</template>

<script>
export default {
  props: {
    parentState: Object
  },

  data () {
    return {
      saving: false,
      specForm: {},
      fieldForm: {},
      fields: [
        { key: 'name', sortable: false },
        { key: 'type', sortable: false },
        { key: 'show_details', label: 'Details', sortable: false }
      ],
      fieldTypes: ['double', 'float', 'int32', 'int64', 'bool', 'string', 'bytes'],
      statuses: {
        draft: 'draft',
        active: 'active',
        archived: 'archived'
      },
      api: {
        baseUrl: '/api/v1/specs'
      }
    }
  },

  watch: {
    'parentState.replicants': function (val) {
    }
  },

  created () {
    this.resetSpecForm()
  },

  mounted () {
  },

  methods: {
    resetSpecForm () {
      this.specForm = {
        name: '',
        package: '',
        respoitory: '',
        fields: [],
        create: true,
        retrieve: true,
        update: true,
        delete: true,
        record_info: {
          status: this.statuses.draft
        }
      }
    },

    resetFieldForm (event) {
      this.fieldForm = {
        name: '',
        description: '',
        type: null,
        sequence: 0,
        is_list: false,
        is_key: false
      }
    },

    handleFieldOk (event) {
      let form = JSON.parse(JSON.stringify(this.fieldForm))

      // the first 2 are reserved and we want 1 more than the current fields
      form.sequence = this.specForm.fields.length + 3

      this.specForm.fields.push(form)
    },

    newField (event) {
      this.$bvModal.show('newFieldModal')
    },

    saveSpec (event) {
      this.saving = true

      let form = JSON.parse(JSON.stringify(this.specForm))
      form.record_info.status = this.statuses.active

      let req = {
        request_id: this.$uuid.v4(),
        payload: [form]
      }

      this.$http.post(this.api.baseUrl, req).then(res => {
        console.debug('successful save for request_id', res.body.request_id)
        // retrieve the current list and update
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
        }).then(() => {
          this.resetSpecForm()
        })
      }).catch(err => {
        console.error(err)
      }).then(() => {
        this.saving = false
        this.parentState.tabIndex = 0
      })
    }
  }
}
</script>

<style>
</style>

Snapshots are read-only copies of volumes that can be used for cloning new
volumes from. Its is possible to create a snapshot through the [Snapshot Create
API](../api/api.md#create-a-snapshot) or the commandline client.

From the command line client, you can type the following to create a snapshot
from an existing volume:

```
$ heketi-cli snapshot create -volume=<vol_name> [-name=<snap_name>]
```

The new snapshot can be used to create a new volume:

```
$ heketi-cli snapshot clone -from-snap=<snap_name> [-name=<clone_name>]
```

The clones of snapshots are new volumes with the same proporties as the
original. The cloned volumes can be deleted with the `heketi-cli volume delete`
command. In a similar fashion, snapshots can be removed with the `heketi-cli
snapshot delete` command.

# Proposed CLI
```
$ heketi-cli snapshot create -volume=<vol_name> [-name=<snap_name>] [-description=<string>]
$ heketi-cli snapshot clone -volume=<vol_name> [-name=<snap_name>]
$ heketi-cli snapshot delete -name=<snap_name>
$ heketi-cli snapshot list [-volume=<vol_name>]
$ heketi-cli snapshot info -name=<snap_name>
```

# API Proposals

## Under the `/volumes` Endpoint

- [x] simple interface, extension to a volume object
- [ ] needs to duplicated/implemented for block-volumes

### Create a Snapshot
* **Method:** _POST_  
* **Endpoint**:`/volumes/{volume_uuid}/snapshots`

### Clone a Snapshot
* **Method:** _POST_  
* **Endpoint**:`/volumes/{volume_uuid}/clone`

### Delete a Snapshot
* **Method:** _DELETE_  
* **Endpoint**:`/volumes/{volume_uuid}/snapshots/{snapshot_uuid}`

### List Snapshots
* **Method:** _GET_  
* **Endpoint**:`/volumes/{volume_uuid}/snapshots`

### Get Snapshot Information
* **Method:** _GET_  
* **Endpoint**:`/volumes/{volume_uuid}/snapshots/{snapshot_uuid}`


## Under the `/snapshots` Endpoint

- [x] simple interface, introduces a new snapshot object
- [x] can be used transparantly for both file+block volumes
- [ ] ugly(?) mixture of file+block volume objects
- [ ] needs to track source (file/block) of the snapshot

### Create a Snapshot
* **Method:** _POST_
* **Endpoint**:`/snapshots/{volume_uuid}`

### Clone a Snapshot
* **Method:** _POST_
* **Endpoint**:`/snapshots/{volume_uuid}/clone`

### Delete a Snapshot
* **Method:** _DELETE_
* **Endpoint**:`/snapshots/{volume_uuid}`

### List Snapshots
* **Method:** _GET_
* **Endpoint**:`/snapshots/{volume_uuid}`

### Get Snapshot Information
* **Method:** _GET_
* **Endpoint**:`/snapshots/{volume_uuid}/{snapshot_uuid}`


# Gluster Snapshot CLI Reference
```
$ gluster --log-file=/dev/null snapshot help

gluster snapshot commands
=========================

snapshot activate <snapname> [force] - Activate snapshot volume.
snapshot clone <clonename> <snapname> - Snapshot Clone.
snapshot config [volname] ([snap-max-hard-limit <count>] [snap-max-soft-limit <percent>]) | ([auto-delete <enable|disable>])| ([activate-on-create <enable|disable>]) - Snapshot Config.
snapshot create <snapname> <volname> [no-timestamp] [description <description>] [force] - Snapshot Create.
snapshot deactivate <snapname> - Deactivate snapshot volume.
snapshot delete (all | snapname | volume <volname>) - Snapshot Delete.
snapshot help - display help for snapshot commands
snapshot info [(snapname | volume <volname>)] - Snapshot Info.
snapshot list [volname] - Snapshot List.
snapshot restore <snapname> - Snapshot Restore.
snapshot status [(snapname | volume <volname>)] - Snapshot Status.
```

- [Snapshot](https://github.com/gluster/glusterfs-specs/blob/master/done/GlusterFS%203.6/Gluster%20Volume%20Snapshot.md)
- [Cloning](https://github.com/gluster/glusterfs-specs/blob/master/done/GlusterFS%203.7/Clone%20of%20Snapshot.md)

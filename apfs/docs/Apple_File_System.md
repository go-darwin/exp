# Apple File System (APFS)

* [Apple File System (APFS)](#apple-file-system-apfs)
  * [Summary](#summary)
  * [Document information](#document-information)
  * [License](#license)
  * [Revision history](#revision-history)
  * [Overview](#overview)
  * [Test version](#test-version)
  * [Terminology](#terminology)
  * [Keys](#keys)
  * [Volume master key](#volume-master-key)
  * [Volume key](#volume-key)
  * [Encryption methods](#encryption-methods)
  * [AES-XTS](#aes-xts)
* [Key bag entries](#key-bag-entries)
  * [Objects](#objects)
  * [Object header](#object-header)
  * [Object checksum](#object-checksum)
  * [Object identifiers](#object-identifiers)
  * [Object types](#object-types)
  * [Object subtypes](#object-subtypes)
  * [B-tree](#b-tree)
  * [B-tree object](#b-tree-object)
     * [B-tree root object](#b-tree-root-object)
     * [B-tree node object](#b-tree-node-object)
  * [B-tree node header](#b-tree-node-header)
     * [<span id="user-content-btree_node_flags"></span>B-tree node flags](#b-tree-node-flags)
  * [B-tree entries](#b-tree-entries)
     * [Fixed-size B-tree entry](#fixed-size-b-tree-entry)
     * [Variable-size B-tree entry](#variable-size-b-tree-entry)
  * [B-tree footer](#b-tree-footer)
     * [<span id="user-content-btree_flags"></span>B-tree flags](#b-tree-flags)
  * [The container](#the-container)
  * [Container superblock](#container-superblock)
     * [<span id="user-content-container_flags"></span>Container flags](#container-flags)
     * [<span id="user-content-container_feature_flags"></span>Container feature flags](#container-feature-flags)
     * [<span id="user-content-container_read_only_feature_flags"></span>Container read-only feature flags](#container-read-only-feature-flags)
     * [<span id="user-content-container_incompatible_feature_flags"></span>Container incompatible feature flags](#container-incompatible-feature-flags)
     * [<span id="user-content-container_counters"></span>Container counters](#container-counters)
     * [Notes](#notes)
  * [Checkpoint map](#checkpoint-map)
     * [Checkpoint map object](#checkpoint-map-object)
     * [<span id="user-content-checkpoint_flags"></span>Checkpoint flags](#checkpoint-flags)
     * [<span id="user-content-checkpoint_map_entry"></span>Checkpoint map entry](#checkpoint-map-entry)
  * [Object map](#object-map)
  * [Object map object](#object-map-object)
     * [<span id="user-content-object_map_flags"></span>Object map flags](#object-map-flags)
  * [Object map B-tree](#object-map-b-tree)
     * [Object map B-tree key](#object-map-b-tree-key)
     * [Object map B-tree branch node value](#object-map-b-tree-branch-node-value)
     * [Object map value](#object-map-value)
        * [<span id="user-content-object_map_value_flags"></span>Object map value flags](#object-map-value-flags)
     * [Notes](#notes-1)
  * [Space manager](#space-manager)
     * [<span id="user-content-space_manager_flags"></span>Space manager flags](#space-manager-flags)
     * [<span id="user-content-space_manager_device"></span>Space manager device](#space-manager-device)
     * [<span id="user-content-space_manager_free_queue"></span>Space manager free queue](#space-manager-free-queue)
     * [<span id="user-content-space_manager_allocation_zone"></span>Space manager allocation zone](#space-manager-allocation-zone)
     * [<span id="user-content-space_manager_zone_boundaries"></span>Space manager zone_boundaries](#space-manager-zone_boundaries)
     * [Notes](#notes-2)
  * [Chunk information address block](#chunk-information-address-block)
  * [Chunk information block](#chunk-information-block)
     * [<span id="user-content-chunk_information_entry"></span>Chunk information entry](#chunk-information-entry)
  * [Reaper](#reaper)
     * [<span id="user-content-reaper_list"></span>Reaper list](#reaper-list)
     * [<span id="user-content-reaper_list_entry"></span>Reaper list entry](#reaper-list-entry)
  * [Key bag](#key-bag)
  * [Container key bag object](#container-key-bag-object)
  * [Volume key bag object](#volume-key-bag-object)
  * [Key bag header](#key-bag-header)
  * [Key bag entries](#key-bag-entries-1)
     * [Key bag entry header](#key-bag-entry-header)
     * [<span id="user-content-key_bag_entry_types"></span>Key bag entry types](#key-bag-entry-types)
        * [Container key bag entry types](#container-key-bag-entry-types)
        * [Volume key bag entry types](#volume-key-bag-entry-types)
     * [<span id="user-content-key_bag_packed_object"></span>Key bag packed object](#key-bag-packed-object)
        * [Key bag packed value](#key-bag-packed-value)
        * [<span id="user-content-key_bag_kek_packed_object"></span>Key encrypted key (KEK) packed object](#key-encrypted-key-kek-packed-object)
        * [<span id="user-content-key_bag_wrapped_kek_packed_object"></span>Wrapped Key Encrypted Key (KEK) packed object](#wrapped-key-encrypted-key-kek-packed-object)
     * [<span id="user-content-wrapped_kek_metadata"></span>Wrapped Key Encrypted Key (KEK) metadata](#wrapped-key-encrypted-key-kek-metadata)
        * [<span id="user-content-encryption_methods"></span>Encryption methods](#encryption-methods-1)
     * [<span id="user-content-key_bag_data_extent"></span>Key bag data extent](#key-bag-data-extent)
  * [Volume](#volume)
  * [Volume superblock](#volume-superblock)
  * [Encryption state](#encryption-state)
     * [<span id="user-content-encryption_state_flags"></span>Encryption state flags](#encryption-state-flags)
  * [Change information](#change-information)
     * [<span id="user-content-volume_superblock_flags"></span>Volume flags](#volume-flags)
     * [<span id="user-content-volume_superblock_feature"></span>Volume superblock features](#volume-superblock-features)
     * [<span id="user-content-volume_read_only_feature_flags"></span>Volume read-only feature flags](#volume-read-only-feature-flags)
     * [<span id="user-content-volume_incompatible_feature_flags"></span>Volume incompatible feature flags](#volume-incompatible-feature-flags)
  * [File system](#file-system)
  * [File system B-tree key](#file-system-b-tree-key)
  * [File system data types](#file-system-data-types)
  * [File system B-tree branch node value](#file-system-b-tree-branch-node-value)
  * [Snapshot metadata](#snapshot-metadata)
     * [<span id="user-content-snapshot_metadata_flags"></span>Snapshot metadata flags](#snapshot-metadata-flags)
  * [Extent](#extent)
     * [Extent key data](#extent-key-data)
     * [Extent value data](#extent-value-data)
  * [Inode](#inode)
     * [Inode key data](#inode-key-data)
     * [Inode value data](#inode-value-data)
        * [<span id="user-content-inode_flags"></span>Inode flags](#inode-flags)
        * [<span id="user-content-file_modes"></span>File modes](#file-modes)
        * [<span id="user-content-bsd_file_entry_flags"></span>BSD file entry flags](#bsd-file-entry-flags)
  * [Extended attribute](#extended-attribute)
     * [Extended attribute key data](#extended-attribute-key-data)
     * [Extended attribute value data](#extended-attribute-value-data)
     * [<span id="user-content-extended_attribute_names"></span>Extended attribute names](#extended-attribute-names)
     * [<span id="user-content-extended_attribute_flags"></span>Extended attribute flags](#extended-attribute-flags)
     * [<span id="user-content-extended_attribute_data_stream"></span>Extended attribute data stream](#extended-attribute-data-stream)
     * [<span id="user-content-compressed_data_extended_attribute"></span>Compressed data extended attribute](#compressed-data-extended-attribute)
        * [<span id="user-content-compressed_data_header"></span>Compressed data header](#compressed-data-header)
        * [<span id="user-content-compression_method"></span>Compression method](#compression-method)
  * [Sibling link](#sibling-link)
     * [Sibling link key data](#sibling-link-key-data)
     * [Sibling link value data](#sibling-link-value-data)
  * [Data stream identifier](#data-stream-identifier)
     * [Data stream identifier key data](#data-stream-identifier-key-data)
     * [Data stream identifier value data](#data-stream-identifier-value-data)
  * [Encryption state](#encryption-state-1)
  * [File extent](#file-extent)
     * [File extent key data](#file-extent-key-data)
     * [File extent value data](#file-extent-value-data)
     * [<span id="user-content-file_extent_flags"></span>File extent flags](#file-extent-flags)
  * [Directory record](#directory-record)
     * [Directory record key data with name](#directory-record-key-data-with-name)
     * [Directory record key data with name and hash](#directory-record-key-data-with-name-and-hash)
     * [Directory record value data](#directory-record-value-data)
        * [<span id="user-content-directory_entry_flags"></span>Directory entry flags](#directory-entry-flags)
        * [<span id="user-content-directory_entry_name_hash"></span>Directory entry name hash](#directory-entry-name-hash)
  * [Directory stats](#directory-stats)
     * [Directory stats key data](#directory-stats-key-data)
     * [Directory stats value data](#directory-stats-value-data)
  * [Snapshot name](#snapshot-name)
  * [Sibling map](#sibling-map)
     * [Sibling map key data](#sibling-map-key-data)
     * [Sibling map value data](#sibling-map-value-data)
  * [Extended fields](#extended-fields)
     * [<span id="user-content-extended_field_descriptor"></span>Extended field descriptor](#extended-field-descriptor)
     * [<span id="user-content-extended_field_types"></span>Extended field types](#extended-field-types)
        * [Directory record extended field types](#directory-record-extended-field-types)
        * [<span id="user-content-inode_extended_field_types"></span>Inode extended field types](#inode-extended-field-types)
     * [<span id="user-content-extended_field_flags"></span>Extended field flags](#extended-field-flags)
  * [Data stream attribute](#data-stream-attribute)
  * [File content](#file-content)
  * [Data fork](#data-fork)
  * [Compressed data extended attribute](#compressed-data-extended-attribute-1)
  * [Compressed data extended attribute with resource fork](#compressed-data-extended-attribute-with-resource-fork)
     * [ZLIB (DEFLATE) compressed data](#zlib-deflate-compressed-data)
        * [ZLIB (DEFLATE) compressed header](#zlib-deflate-compressed-header)
        * [ZLIB (DEFLATE) compressed data block descriptors](#zlib-deflate-compressed-data-block-descriptors)
        * [ZLIB (DEFLATE) compressed data block descriptor](#zlib-deflate-compressed-data-block-descriptor)
        * [ZLIB (DEFLATE) compressed footer](#zlib-deflate-compressed-footer)
     * [LZVN compressed data](#lzvn-compressed-data)
  * [Resource fork](#resource-fork)
  * [Extended attribute (named fork)](#extended-attribute-named-fork)
  * [EFI jumpstart](#efi-jumpstart)
  * [EFI jumpstart extent](#efi-jumpstart-extent)
  * [Extent-reference tree](#extent-reference-tree)
  * [Snapshots](#snapshots)
  * [Snapshot metadata tree](#snapshot-metadata-tree)
  * [Snapshot metadata tree object](#snapshot-metadata-tree-object)
  * [Snapshot metadata B-tree](#snapshot-metadata-b-tree)
     * [Snapshot metadata B-tree key](#snapshot-metadata-b-tree-key)
     * [Snapshot metadata B-tree branch node value](#snapshot-metadata-b-tree-branch-node-value)
     * [Snapshot metadata B-tree leaf node value](#snapshot-metadata-b-tree-leaf-node-value)
  * [Fusion drives](#fusion-drives)
  * [Fusion middle tree](#fusion-middle-tree)
  * [Notes](#notes-3)
  * [References](#references)

## Summary

The Apple File System (APFS) is a volume and file system mainly used
on the Apple platforms such as MacOS and iOS. APFS succeeds the
Hierarchical File System (HFS) and Core Storage (CS). This
specification is based on publicly available work on the format and
was enhanced by analyzing test data.

This document is intended as a working document of the data format
specification for the libfsapfs project.

## Document information

<table>
<colgroup>
<col style="width: 16%" />
<col style="width: 83%" />
</colgroup>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>Author(s):</p></td>
<td style="text-align: left;"><p>Joachim Metz &lt;<a href="mailto:joachim.metz@gmail.com">joachim.metz@gmail.com</a>&gt;</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Abstract:</p></td>
<td style="text-align: left;"><p>This document contains information about the Apple File System (APFS).</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>Classification:</p></td>
<td style="text-align: left;"><p>Public</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Keywords:</p></td>
<td style="text-align: left;"><p>Apple File System, APFS</p></td>
</tr>
</tbody>
</table>

## License

    Copyright (C) 2018-2020, Joachim Metz <joachim.metz@gmail.com>.
    Permission is granted to copy, distribute and/or modify this document under the
    terms of the GNU Free Documentation License, Version 1.3 or any later version
    published by the Free Software Foundation; with no Invariant Sections, no
    Front-Cover Texts, and no Back-Cover Texts. A copy of the license is included
    in the section entitled "GNU Free Documentation License".

## Revision history

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Version</th>
<th style="text-align: left;">Author</th>
<th style="text-align: left;">Date</th>
<th style="text-align: left;">Comments</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0.0.1</p></td>
<td style="text-align: left;"><p>J.B. Metz</p></td>
<td style="text-align: left;"><p>September 2018</p></td>
<td style="text-align: left;"><p>Initial version.</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0.0.2</p></td>
<td style="text-align: left;"><p>J.B. Metz</p></td>
<td style="text-align: left;"><p>October 2018</p></td>
<td style="text-align: left;"><p>Additional format information.</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0.0.3</p></td>
<td style="text-align: left;"><p>J.B. Metz</p></td>
<td style="text-align: left;"><p>November 2018</p></td>
<td style="text-align: left;"><p>Additional format information.</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0.0.4</p></td>
<td style="text-align: left;"><p>J.B. Metz</p></td>
<td style="text-align: left;"><p>December 2018</p></td>
<td style="text-align: left;"><p>Additional format information.</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0.0.5</p></td>
<td style="text-align: left;"><p>J.B. Metz</p></td>
<td style="text-align: left;"><p>January 2019</p></td>
<td style="text-align: left;"><p>Additional format information.</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0.0.6</p></td>
<td style="text-align: left;"><p>J.B. Metz</p></td>
<td style="text-align: left;"><p>May 2019</p></td>
<td style="text-align: left;"><p>Additional format information.</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0.0.7</p></td>
<td style="text-align: left;"><p>J.B. Metz</p></td>
<td style="text-align: left;"><p>June 2019</p></td>
<td style="text-align: left;"><p>Additional format information.</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0.0.8</p></td>
<td style="text-align: left;"><p>J.B. Metz</p></td>
<td style="text-align: left;"><p>September 2019</p></td>
<td style="text-align: left;"><p>Additional format information regarding container key bag.</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0.0.9</p></td>
<td style="text-align: left;"><p>J.B. Metz</p></td>
<td style="text-align: left;"><p>October 2020</p></td>
<td style="text-align: left;"><p>Additional format information regarding embedded extended attributes.</p></td>
</tr>
</tbody>
</table>

## Overview

An APFS volume consists of:

-   Container

    -   Zero or more volumes

<table>
<colgroup>
<col style="width: 16%" />
<col style="width: 83%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Characteristics</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>Byte order</p></td>
<td style="text-align: left;"><p>little-endian</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Date and time values</p></td>
<td style="text-align: left;"><p>number of nanoseconds since January 1, 1970 00:00:00 UTC (POSIX epoch), disregarding leap seconds</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>Character strings</p></td>
<td style="text-align: left;"><p>Unicode strings are stored in UTF-8</p></td>
</tr>
</tbody>
</table>

The date and values are signed to represent dates before January 1,
1970. Other sources are known to claim the date and time values are
unsigned including Apple’s reference documentation `[APPLE18]`.

Although (some) Apple sources claim that APFS uses Unicode version 9.0,
support for more recent Unicode versions has been observed.

## Test version

The following version of programs were used to test the information
within this document:

-   macOS 10.13 (High Sierra)

-   macOS 10.14 (Mojave)

## Terminology

"Physical" volume, the volume in which the APFS container is stored.

"Logical" volume, the volume in which an APFS file system is stored.

## Keys

To encrypt storage media APFS uses different kind of keys.

## Volume master key

The Volume Master Key (VMK) is used to encrypt the data of a specific
volume.

**<span class="yellow-background">TODO: complete this section.</span>**

## Volume key

For every volume on an Mac OS system with APFS, APFS provides for a
volume password to unlock the encrypted data. The volume password is
used to determine a volume key.

**<span class="yellow-background">TODO: complete this section.</span>**

## Encryption methods

APFS uses the AES-XTS encryption method to encrypt the key bag, file
system metadata and content data.

## AES-XTS

The AES-XTS encryption method uses:

-   a primary key (key 1) to encrypt/decrypt the data (the whitened
    plaintext/ciphertext).

-   a secondary key (key 2) to encrypt/ decrypt the tweak value, also
    referred to as the tweak key. The encrypted tweak value is used to
    whiten the plaintext/ciphertext.

-   a tweak value

The cipher block size is 128 bytes.

See `[IEEE 1619-2007]` for more information.

The container key bag is encrypted using the "container identifier" of
the container as both the primary and tweak key. The sector number,
relative to the start of the container, is used as the tweak value.

When a T2 chip is present, it is currently assumed that the T2 is used
to encrypte the container key bag instead of the "container identifier".

The unit size is the sector size, which is assumed to be 512 bytes also
for 4k sector media.

The volume key bag is encrypted using the "volume identifier" of the
corresponding key bag entry, as both the primary and tweak key. The
sector number, relative to the start of the container, is used as the
tweak value.

The file system B-tree is encrypted using the volume master key and the
sector number, relative to the start of the container, is used as the
tweak value.

**<span class="yellow-background">TODO: complete this section.</span>**

# <span id="key_bag_entries"></span>Key bag entries

## Objects

APFS uses the "object" data type to distinguish between different data
types.

## Object header

The object header (obj\_phys\_t) is 32 bytes of size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Checksum (o_cksum)<br />
See section: <a href="#object_checksum">Object checkum</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object identifier (o_oid)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object transaction identifier (o_xid)<br />
Identifier of the most recent transaction that this object was modified in</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object type (o_type)<br />
See section: <a href="#object_types">Object types</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>28</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object subtype (o_subtype)<br />
See section: <a href="#object_subtypes">Object subtypes</a></p></td>
</tr>
</tbody>
</table>

## <span id="object_checksum"></span>Object checksum

The checksum algorithm:

-   calculate a Fletcher-64 checksum of the block data without the
    object checkum value and an initial value of 0

-   `checksum_lower_32bit = (fletcher_lower_32bit + fletcher_upper_32bit) mod 0xffffffff`

-   `checksum_upper_32bit = (fletcher_lower_32bit + checksum_lower_32bit) mod 0xffffffff`

-   `checksum = (checksum_upper_32bit << 32) | checksum_lower_32bit`

## Object identifiers

-   For a physical object, its identifier is the logical block address
    on disk where the object is stored.

-   For an ephemeral object, its identifier is a number.

-   For a virtual object, its identifier is a number.

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>OID_INVALID</p></td>
<td style="text-align: left;"><p>Invalid</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>1</p></td>
<td style="text-align: left;"><p>OID_NX_SUPERBLOCK</p></td>
<td style="text-align: left;"><p>Container superblock</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>1024</p></td>
<td style="text-align: left;"><p>OID_RESERVED_COUNT</p></td>
<td style="text-align: left;"><p>Number of reserved object identifiers</p></td>
</tr>
</tbody>
</table>

## <span id="object_types"></span>Object types

The object type (o\_type) value consists of a type and flags.

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0x00000000</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_INVALID</p></td>
<td style="text-align: left;"><p>Invalid<br />
For a subtype this value represents not set or not specified</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00000001</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_NX_SUPERBLOCK</p></td>
<td style="text-align: left;"><p>Container superblock<br />
See section: <a href="#container_superblock">Container superblock</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00000002</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_BTREE</p></td>
<td style="text-align: left;"><p>B-Tree (root)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00000003</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_BTREE_NODE</p></td>
<td style="text-align: left;"><p>B-Tree node</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00000004</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (MTree?)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00000005</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_SPACEMAN</p></td>
<td style="text-align: left;"><p>Space manager header</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00000006</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_SPACEMAN_CAB</p></td>
<td style="text-align: left;"><p>Space manager chunk information address block<br />
See section: <a href="#chunk_information_address_block">Chunk information address block</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00000007</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_SPACEMAN_CIB</p></td>
<td style="text-align: left;"><p>Space manager chunk information block<br />
See section: <a href="#chunk_information_block">Chunk information block</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00000008</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_SPACEMAN_BITMAP</p></td>
<td style="text-align: left;"><p>Space manager bitmap</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00000009</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_SPACEMAN_FREE_QUEUE</p></td>
<td style="text-align: left;"><p>Space manager free queue</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x0000000a</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_EXTENT_LIST_TREE</p></td>
<td style="text-align: left;"><p>Extent list tree</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0000000b</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_OMAP</p></td>
<td style="text-align: left;"><p>Object map<br />
See section: <a href="#object_map">Object map</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x0000000c</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_CHECKPOINT_MAP</p></td>
<td style="text-align: left;"><p>Checkpoint map</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0000000d</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_FS</p></td>
<td style="text-align: left;"><p>Volume superblock (File system)<br />
See section: <a href="#volume_superblock">Volume superblock</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x0000000e</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_FS</p></td>
<td style="text-align: left;"><p>File system tree<br />
See section: <a href="#file_system">File system</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0000000f</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_BLOCKREFTREE</p></td>
<td style="text-align: left;"><p>Extent-reference tree<br />
See section: <a href="#extent_reference_tree">Extent reference tree</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00000010</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_SNAPMETATREE</p></td>
<td style="text-align: left;"><p>Snapshot metadata tree<br />
See section: <a href="#snapshot_metadata_tree">Snapshot metadata tree</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00000011</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_NX_REAPER</p></td>
<td style="text-align: left;"><p>Reaper<br />
See section: <a href="#reaper">Reaper</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00000012</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_NX_REAP_LIST</p></td>
<td style="text-align: left;"><p>Reaper list<br />
See section: <a href="#reaper_list">Reaper list</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00000013</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_OMAP_SNAPSHOT</p></td>
<td style="text-align: left;"><p>Object map snapshot</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00000014</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_EFI_JUMPSTART</p></td>
<td style="text-align: left;"><p>EFI jumpstart<br />
See section: <a href="#efi_jumpstart">EFI jumpstart</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00000015</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_FUSION_MIDDLE_TREE</p></td>
<td style="text-align: left;"><p>Fusion middle tree<br />
See section: <a href="#fusion_middle_tree">Fusion middle tree</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00000016</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_NX_FUSION_WBC</p></td>
<td style="text-align: left;"><p>Fusion write-back cache<br />
See section: <a href="#fusion_write_back_cache">Fusion write-back cache</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00000017</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_NX_FUSION_WBC_LIST</p></td>
<td style="text-align: left;"><p>Fusion write-back cache list<br />
See section: <a href="#fusion_write_back_cache">Fusion write-back cache</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00000018</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_ER_STATE</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (ER state?)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00000019</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_GBITMAP</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (G Bitmap?)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x0000001a</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_GBITMAP_TREE</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (G Bitmap tree?)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0000001b</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_GBITMAP_BLOCK</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (G Bitmap block?)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x000000ff</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_TEST</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (test?)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0000ffff</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_MASK</p></td>
<td style="text-align: left;"><p>Object type bitmask</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;" colspan="3"><p><em>Flags used in combination with some of the object types</em></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x08000000</p></td>
<td style="text-align: left;"><p>OBJ_NONPERSISTENT</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (Non-persistent?)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x10000000</p></td>
<td style="text-align: left;"><p>OBJ_ENCRYPTED</p></td>
<td style="text-align: left;"><p>Is encrypted</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x20000000</p></td>
<td style="text-align: left;"><p>OBJ_NOHEADER</p></td>
<td style="text-align: left;"><p>Has no object (obj_phys_t) header</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00000000</p></td>
<td style="text-align: left;"><p>OBJ_VIRTUAL</p></td>
<td style="text-align: left;"><p>Is virtual object</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x40000000</p></td>
<td style="text-align: left;"><p>OBJ_PHYSICAL</p></td>
<td style="text-align: left;"><p>Is physical object</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x80000000</p></td>
<td style="text-align: left;"><p>OBJ_EPHEMERAL</p></td>
<td style="text-align: left;"><p>Is ephemeral object</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0xffff0000</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_FLAGS_MASK</p></td>
<td style="text-align: left;"><p>Object type flags bitmask</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0xc0000000</p></td>
<td style="text-align: left;"><p>OBJ_STORAGETYPE_MASK</p></td>
<td style="text-align: left;"><p>Object storage type bitmast</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0xf8000000</p></td>
<td style="text-align: left;"><p>OBJECT_TYPE_FLAGS_DEFINED_MASK</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;" colspan="3"><p><em>Object types without flags</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x6b657973</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Container key bag</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x72656373</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Volume key bag</p></td>
</tr>
</tbody>
</table>

## <span id="object_subtypes"></span>Object subtypes

The object subtype is used by specific object types such as:

-   B-Tree root

-   B-Tree node

The object subtypes are the same as the [Object types](#object_types).

## <span id="btree"></span>B-tree

A B-tree consists of:

-   B-tree (root or node) object

-   B-tree node header

-   B-tree entries (table of contents)

-   keys data, where the first key is stored after the entries in
    increasing order

-   Optional key free list

-   unused data

-   Optional value free list

-   values data, where the first value is stored before the footer in
    descending order

-   Optional B-tree footer, which is only stored in the root node

`[APPLE18]` combines the B-Tree object and B-tree node header into a
single structure referred to as btree\_node\_phys\_t.

## B-tree object

### B-tree root object

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object header (btn_o)</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object checksum<br />
See section: <a href="#object_checksum">Object checkum</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object identifier</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object transaction identifier (xid)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x00000002<br />
0x40000002</p></td>
<td style="text-align: left;"><p>Object type<br />
See section: <a href="#object_types">Object types</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>28</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object subtype<br />
See section: <a href="#object_subtypes">Object subtypes</a></p></td>
</tr>
</tbody>
</table>

### B-tree node object

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object header (btn_o)</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object checksum<br />
See section: <a href="#object_checksum">Object checkum</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object identifier</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object transaction identifier (xid)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x00000003<br />
0x40000003</p></td>
<td style="text-align: left;"><p>Object type<br />
See section: <a href="#object_types">Object types</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>28</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object subtype<br />
See section: <a href="#object_subtypes">Object subtypes</a></p></td>
</tr>
</tbody>
</table>

## B-tree node header

The B-tree node header is stored after the B-tree root or node object.

The B-tree node header is 24 bytes of size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Flags (btn_flags)<br />
See section: <a href="#btree_node_flags">B-tree node flags</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Level (btn_level)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of keys in the node (btn_nkeys)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;" colspan="4"><p><em>Table space (btn_table_space)</em></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Entries data offset<br />
Contains an offset relative to the end of the B-tree node header or -1 (0xffff) if not set (invalid)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>10</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Entries data size</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Free space (btn_free_space)</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>12</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Unused data offset<br />
Contains an offset relative to the end of the entries data or -1 (0xffff) if not set (invalid)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>14</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Unused data size</p></td>
</tr>
<tr class="even">
<td style="text-align: left;" colspan="4"><p><em>Key free list (btn_key_free_list)</em></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Unused key list offset<br />
Contains an offset relative to ??? or -1 (0xffff) if not set (invalid)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>18</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Unused key list size</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Value free list (btn_val_free_list)</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>20</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Unused value list offset<br />
Contains an offset relative to ??? or -1 (0xffff) if not set (invalid)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>22</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Unused value list size</p></td>
</tr>
</tbody>
</table>

### <span id="btree_node_flags"></span>B-tree node flags

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0x0001</p></td>
<td style="text-align: left;"><p>BTNODE_ROOT</p></td>
<td style="text-align: left;"><p>Is root</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0002</p></td>
<td style="text-align: left;"><p>BTNODE_LEAF</p></td>
<td style="text-align: left;"><p>Is leaf</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x0004</p></td>
<td style="text-align: left;"><p>BTNODE_FIXED_KV_SIZE</p></td>
<td style="text-align: left;"><p>Has a fixed-size entry (key and value)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x8000</p></td>
<td style="text-align: left;"><p>BTNODE_CHECK_KOFF_INVAL</p></td>
<td style="text-align: left;"><p>In transient state<br />
This flag is used for in-memory purposes only</p></td>
</tr>
</tbody>
</table>

## B-tree entries

The B-tree entries are stored after the B-tree node header.

### Fixed-size B-tree entry

The fixed-size B-tree entry is 4 bytes of size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Key data offset (key_offs)<br />
Contains an offset relative to the end of the entries data</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Value data offset (value_offs)<br />
Contains a reversed offset relative to the start of the B-Tree footer</p></td>
</tr>
</tbody>
</table>

### Variable-size B-tree entry

The variable-size B-tree entry is 8 bytes of size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Key data offset (key_offs)<br />
Contains an offset relative to the end of the entries data</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Key data size (key_len)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Value data offset (value_offs)<br />
Contains a reversed offset relative to the start of the B-Tree footer</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>6</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Value data size (value_len)</p></td>
</tr>
</tbody>
</table>

## B-tree footer

The B-tree footer is stored at the end of the block that contains the
B-tree root boject.

The B-tree footer (btree\_info\_t) is 40 bytes of size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Static information (btree_info_fixed_t)</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Flags (bt_flags)<br />
See section: <a href="#btree_flags">B-tree flags</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Node size (bt_node_size)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Key size (bt_key_size)<br />
Set to 0 if key has a variable size</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>12</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Value size (bt_val_size)<br />
Set to 0 if value has a variable size</p></td>
</tr>
<tr class="even">
<td style="text-align: left;" colspan="4"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Maximum key size (bt_longest_key)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>20</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Maximum value size (bt_longest_val)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Total number of keys (bt_key_count)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>32</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Total number of nodes (bt_node_count)</p></td>
</tr>
</tbody>
</table>

### <span id="btree_flags"></span>B-tree flags

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0x00000001</p></td>
<td style="text-align: left;"><p>BTREE_UINT64_KEYS</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00000002</p></td>
<td style="text-align: left;"><p>BTREE_SEQUENTIAL_INSERT</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00000004</p></td>
<td style="text-align: left;"><p>BTREE_ALLOW_GHOSTS</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00000008</p></td>
<td style="text-align: left;"><p>BTREE_EPHEMERAL</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00000010</p></td>
<td style="text-align: left;"><p>BTREE_PHYSICAL</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00000020</p></td>
<td style="text-align: left;"><p>BTREE_NONPERSISTENT</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00000040</p></td>
<td style="text-align: left;"><p>BTREE_KV_NONALIGNED</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
</tbody>
</table>

## The container

APFS stores volumes inside a container. The maximum number of volumes is
dependent on the size of the container. `[HANSEN17]` indicates:

<table>
<colgroup>
<col style="width: 16%" />
<col style="width: 83%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Container size</th>
<th style="text-align: left;">Maximum number of volumes</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>1 GiB</p></td>
<td style="text-align: left;"><p>2</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>2 GiB</p></td>
<td style="text-align: left;"><p>4</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>5 GiB</p></td>
<td style="text-align: left;"><p>10</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>10 GiB</p></td>
<td style="text-align: left;"><p>20</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>20 GiB</p></td>
<td style="text-align: left;"><p>40</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>100 GiB</p></td>
<td style="text-align: left;"><p>100</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>12 TiB</p></td>
<td style="text-align: left;"><p>100</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>1.2 PiB</p></td>
<td style="text-align: left;"><p>100</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>7.5 EiB</p></td>
<td style="text-align: left;"><p>100</p></td>
</tr>
</tbody>
</table>

The container consists of:

-   current container superblock

-   stored in the container checkpoint descriptor area:

    -   current checkpoint map

    -   previous checkpoint map(s)

    -   previous container superblock(s)

-   stored in the container:

    -   space manager

    -   container object map

    -   reaper

    -   crypto key

    -   zero or more volumes

-   **<span class="yellow-background">backup of current container
    superblock?</span>**

## <span id="container_superblock"></span>Container superblock

The container superblock is 1382 bytes of size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object header</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object checksum<br />
See section: <a href="#object_checksum">Object checkum</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object identifier</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object transaction identifier (xid)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x80000001</p></td>
<td style="text-align: left;"><p>Object type<br />
See section: <a href="#object_types">Object types</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>28</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x00000000</p></td>
<td style="text-align: left;"><p>Object subtype</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object values</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>32</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>"NXSB"</p></td>
<td style="text-align: left;"><p>Signature (nx_magix)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>36</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Block size (nx_block_size)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>40</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of blocks (nx_block_count)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>48</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Compatible feature flags (nx_features)<br />
See section: <a href="#container_feature_flags">Container feature flags</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>56</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Read-only compatible feature flags (nx_readonly_compatible_features)<br />
See section: <a href="#container_read_only_feature_flags">Container read-only feature flags</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>64</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Incompatible feature flags (nx_incompatible_features)<br />
See section: <a href="#container_incompatible_feature_flags">Container incompatible feature flags</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>72</p></td>
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Container identifier (nx_uuid)<br />
Contains a UUID stored in big-endian</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>88</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Next (available) object identifier (nx_next_oid)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>96</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Next (available) transaction identifier (nx_next_xid)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>104</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Checkpoint descriptor area number of blocks (nx_xp_desc_blocks)<br />
Contains size, in the number of blocks, of the checkpoint descriptor area and the MSB is a flag</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>108</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Checkpoint data area number of blocks (nx_xp_data_blocks)<br />
Contains size, in the number of blocks, of the checkpoint data area and the MSB is a flag</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>112</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Checkpoint descriptor area block number (nx_xp_desc_base)<br />
Contains the block number relative to the start of the container of the checkpoint descriptor area if the MSB of nx_xp_desc_blocks is not set otherwise the value contains the physical object identifier of a checkpoint descriptor area B-tree</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>120</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Checkpoint data area block number (nx_xp_data_base)<br />
Contains the block number relative to the start of the container of the checkpoint data area if the MSB of nx_xp_data_blocks is not set</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>128</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Next available index in the checkpoint descriptor area (nx_xp_desc_next)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>132</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Next available index in the checkpoint data area (nx_xp_data_next)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>136</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Start index in the checkpoint descriptor area used by the superblock (nx_xp_desc_index)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>140</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of blocks in the checkpoint descriptor area used by the superblock (nx_xp_desc_len)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>144</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Start index in the checkpoint data area used by the superblock (nx_xp_data_index)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>148</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of blocks in the checkpoint data area used by the superblock (nx_xp_data_len)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>152</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Space manager object identifier (nx_spaceman_oid)<br />
Contains a object identifier that can be resolved in the <a href="#checkpoint_map">checkpoint map</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>160</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object map block number (nx_omap_oid)<br />
Contains a block number relative to the start of the container of the <a href="#object_map">object map</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>168</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Reaper object identifier (nx_reaper_oid)<br />
Contains a object identifier that can be resolved in the <a href="#checkpoint_map">checkpoint map</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>176</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (nx_test_type)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>180</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Maxmum number of volumes (nx_max_file_systems)<br />
Contains the maximum number of volumes supported by the container</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>184</p></td>
<td style="text-align: left;"><p>100 x 8 = 800</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Array of volume object identifiers (nx_fs_oid)<br />
The object identifiers can be resolved in the <a href="#object_map">object map</a> to a "physical" location</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>984</p></td>
<td style="text-align: left;"><p>32 x 8 = 256</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Counters (nx_counters)<br />
See section: <a href="#container_counters">Container counters</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;" colspan="4"><p><em>Unknown (nx_blocked_out_prange)</em></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>1240</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (nx_blocked_out_base)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>1248</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (nx_blocked_out_blocks)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;" colspan="4"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>1254</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (nx_evict_mapping_tree_oid)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>1262</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Container flags (nx_flags)<br />
See section: <a href="#container_flags">Container flags</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>1270</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>EFI jumpstart (physical) object identifier (nx_efi_jumpstart)<br />
Contains a block number relative to the start of the container of the <a href="#efi_jumpstart">EFI jumpstart</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>1278</p></td>
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Fusion set identifier (nx_fusion_uuid)<br />
Contains a UUID stored in big-endian</p></td>
</tr>
<tr class="even">
<td style="text-align: left;" colspan="4"><p><em>Unknown (nx_keylocker)</em></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>1294</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Container key bag data block number (nx_keybag_base)<br />
Contains a block number relative to the start of the container of the <a href="#key_bag">Key bag</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>1302</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Contaner key bag data number of blocks (nx_keybag_blocks)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;" colspan="4"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>1310</p></td>
<td style="text-align: left;"><p>4 x 8 = 32</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (nx_ephemeral_info)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>1342</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (Test object identifier) (nx_test_oid)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>1350</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Fusion middle tree block number (nx_fusion_mt_oid)<br />
Contains a block number relative to the start of the container of the <a href="#fusion_middle_tree">Fusion middle tree</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>1358</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Fusion write-back cache state object identifier (nx_fusion_wbc_oid)<br />
Contains a object identifier that can be resolved in the <a href="#checkpoint_map">checkpoint map</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;" colspan="4"><p><em>Unknown (nx_fusion_wbc)</em></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>1366</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Start block of the Fusion write-back cache area (nx_fusion_wbc_base)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>1374</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of blocks of the Fusion write-back cache area (nx_fusion_wbc_blocks)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;" colspan="4"></td>
</tr>
</tbody>
</table>

Presumably NXSB is an abbreviation of NX superblock. At this point it is
unclear what NX stands for.

### <span id="container_flags"></span>Container flags

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0x00000001</p></td>
<td style="text-align: left;"><p>NX_RESERVED_1</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (reserved)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00000002</p></td>
<td style="text-align: left;"><p>NX_RESERVED_2</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (reserved)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00000004</p></td>
<td style="text-align: left;"><p>NX_CRYPTO_SW</p></td>
<td style="text-align: left;"><p>The encryption is performed in software</p></td>
</tr>
</tbody>
</table>

### <span id="container_feature_flags"></span>Container feature flags

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0x0000000000000001</p></td>
<td style="text-align: left;"><p>NX_FEATURE_DEFRAG</p></td>
<td style="text-align: left;"><p>Supports defragmentation</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0000000000000002</p></td>
<td style="text-align: left;"><p>NX_FEATURE_LCFD</p></td>
<td style="text-align: left;"><p>Use low-capacity Fusion Drive mode</p></td>
</tr>
</tbody>
</table>

### <span id="container_read_only_feature_flags"></span>Container read-only feature flags

Current no read-only feature flags are defined

### <span id="container_incompatible_feature_flags"></span>Container incompatible feature flags

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0x0000000000000001</p></td>
<td style="text-align: left;"><p>NX_INCOMPAT_VERSION1</p></td>
<td style="text-align: left;"><p>Pre-release version 1 of APFS</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0000000000000002</p></td>
<td style="text-align: left;"><p>NX_INCOMPAT_VERSION2</p></td>
<td style="text-align: left;"><p>Release version 2 of APFS</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0000000000000100</p></td>
<td style="text-align: left;"><p>NX_INCOMPAT_FUSION</p></td>
<td style="text-align: left;"><p>Supports Fusion Drives</p></td>
</tr>
</tbody>
</table>

According to `[APPLE18]` the pre-release version 1 and release version 2
are incompatble.

### <span id="container_counters"></span>Container counters

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>NX_CNTR_OBJ_CKSUM_SET</p></td>
<td style="text-align: left;"><p>Number of times a checksum has been calculated when wrting to disk</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>1</p></td>
<td style="text-align: left;"><p>NX_CNTR_OBJ_CKSUM_FAIL</p></td>
<td style="text-align: left;"><p>Number of checksum errors when reading from disk</p></td>
</tr>
</tbody>
</table>

The other 30 counters are presumed to be unused at this point.

### Notes

    checkpoint descriptor area B-tree
    The treeʼs keys are block offsets into the checkpoint descriptor area, and its
    values are instances of prange_t that contain the fragmentʼs size and location.

    checkpoint data area B-tree
    The treeʼs keys are block offsets into the checkpoint data area, and its values are instances of
    prange_t that contain the fragmentʼs size and location.

## <span id="checkpoint_map"></span>Checkpoint map

The checkpoint map contains a mapping between container metadata object
identifiers and their location in the volume.

### Checkpoint map object

The checkpoint map object (checkpoint\_map\_phys\_t) is 4080 bytes of
size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object header</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object checksum<br />
See section: <a href="#object_checksum">Object checkum</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object identifier</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object transaction identifier (xid)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x4000000c</p></td>
<td style="text-align: left;"><p>Object type<br />
See section: <a href="#object_types">Object types</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>28</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x00000000</p></td>
<td style="text-align: left;"><p>Object subtype</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object values</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>32</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Flags (cpm_flags)<br />
See section: <a href="#checkpoint_flags">Checkpoint flags</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>36</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of entries (cpm_count)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>40</p></td>
<td style="text-align: left;"><p>101 x 40 = 4040</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Array of map entries (cpm_map)<br />
See sections: <a href="#checkpoint_map_entry">Checkpoint map entry</a></p></td>
</tr>
</tbody>
</table>

### <span id="checkpoint_flags"></span>Checkpoint flags

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0x00000001</p></td>
<td style="text-align: left;"><p>CHECKPOINT_MAP_LAST</p></td>
<td style="text-align: left;"><p>Last checkpoint map object</p></td>
</tr>
</tbody>
</table>

### <span id="checkpoint_map_entry"></span>Checkpoint map entry

The checkpoint map (checkpoint\_mapping\_t) entry entry is 40 bytes of
size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object header</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>(Container) object type (cpm_type)<br />
See section: <a href="#object_types">Object types</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>(Container) object subtype (cpm_subtype)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Size (cpm_size)<br />
Contains number of bytes</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>12</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Padding (cpm_pad)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>File system object indentifier (cpm_fs_oid)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>(Container) object identifier (cpm_oid)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>32</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Physical address (cpm_paddr)<br />
Contains a block number relative to the start of the container</p></td>
</tr>
</tbody>
</table>

## <span id="object_map"></span>Object map

The object map contains a mapping between object identifiers and their
"physical" location.

The object map consists of:

-   object map (object)

-   object map B-tree

## Object map object

The object map object (omap\_phys\_t) is 88 bytes of size and consists
of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object header</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object checksum<br />
See section: <a href="#object_checksum">Object checkum</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object identifier</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object transaction identifier (xid)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x4000000b</p></td>
<td style="text-align: left;"><p>Object type<br />
See section: <a href="#object_types">Object types</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>28</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x00000000</p></td>
<td style="text-align: left;"><p>Object subtype</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object values</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>32</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Flags (om_flags)<br />
See section: <a href="#object_map_flags">Object map flags</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>36</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of snapshots (om_snap_count)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>40</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object map B-tree type (om_tree_type)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>44</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object map snapshots B-tree type (om_snapshot_tree_type)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>48</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object map B-tree object identifier (om_tree_oid)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>56</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object map snapshots B-tree object identifier (om_snapshot_tree_oid)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>64</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Most recent snapshot object identifier (om_most_recent_snap)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>72</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown transaction identifier (om_pending_revert_min)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>80</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown transaction identifier (om_pending_revert_max)</span></strong></p></td>
</tr>
</tbody>
</table>

### <span id="object_map_flags"></span>Object map flags

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0x00000001</p></td>
<td style="text-align: left;"><p>OMAP_MANUALLY_MANAGED</p></td>
<td style="text-align: left;"><p>No snapshot support</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00000002</p></td>
<td style="text-align: left;"><p>OMAP_ENCRYPTING</p></td>
<td style="text-align: left;"><p>Encryption in progress</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00000004</p></td>
<td style="text-align: left;"><p>OMAP_DECRYPTING</p></td>
<td style="text-align: left;"><p>Decryption in progress</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00000008</p></td>
<td style="text-align: left;"><p>OMAP_KEYROLLING</p></td>
<td style="text-align: left;"><p>Re-encryption with new key in progress</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00000010</p></td>
<td style="text-align: left;"><p>OMAP_CRYPTO_GENERATION</p></td>
<td style="text-align: left;"><p>Encryption configuation has changed</p></td>
</tr>
</tbody>
</table>

## Object map B-tree

The object map values are stored in [B-tree](#btree).

### Object map B-tree key

The object map B-tree key (omap\_key\_t) is 16 bytes of size and
consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Key object identifier (ok_oid)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Key object transaction identifier (ok_xid)</p></td>
</tr>
</tbody>
</table>

### Object map B-tree branch node value

An object map B-tree node contains branch node values if BTNODE\_LEAF is
not set. The corresponding object map B-tree key represents the first
key in the branch.

An object map B-tree branch node value is 8 bytes of size and consists
of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Sub node block number<br />
Contains a block number relative to the start of the container</p></td>
</tr>
</tbody>
</table>

### Object map value

An object map B-tree node contains object map values if BTNODE\_LEAF is
set.

The object map value (omap\_val\_t) is 16 bytes of size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Value object flags (ov_flags)<br />
See section: <a href="#object_map_value_flags">Object map value flags</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Value object size (ov_size)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Value object physical address (ov_paddr)<br />
Contains a block number relative to the start of the container</p></td>
</tr>
</tbody>
</table>

#### <span id="object_map_value_flags"></span>Object map value flags

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0x00000001</p></td>
<td style="text-align: left;"><p>OMAP_VAL_DELETED</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00000002</p></td>
<td style="text-align: left;"><p>OMAP_VAL_SAVED</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00000004</p></td>
<td style="text-align: left;"><p>OMAP_VAL_ENCRYPTED</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00000008</p></td>
<td style="text-align: left;"><p>OMAP_VAL_NOHEADER</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00000010</p></td>
<td style="text-align: left;"><p>OMAP_VAL_CRYPTO_GENERATION</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
</tbody>
</table>

### Notes

TODO document omap\_snapshot\_t TODO document Object Map Reaper Phases

## Space manager

The space manager (spaceman\_phys\_t) is variable of size and consists
of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object header (sm_o)</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object checksum<br />
See section: <a href="#object_checksum">Object checkum</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object identifier</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object transaction identifier (xid)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x80000005</p></td>
<td style="text-align: left;"><p>Object type<br />
See section: <a href="#object_types">Object types</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>28</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x00000000</p></td>
<td style="text-align: left;"><p>Object subtype</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object values</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>32</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Block size (sm_block_size)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>36</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of blocks per chunk (sm_blocks_per_chunk)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>40</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of chunks per chunk information block (CIB) (sm_chunks_per_cib)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>44</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of chunk information blocks (CIBs) per chunk information address block (CAB) (sm_cibs_per_cab)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;" colspan="4"><p><em>Space manager devices (sm_dev)</em></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>48</p></td>
<td style="text-align: left;"><p>48</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Main device (SD_MAIN)<br />
Contains a <a href="#space_manager_device">Space manager device</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>96</p></td>
<td style="text-align: left;"><p>48</p></td>
<td style="text-align: left;"><p>Tier2 device (SD_TIER2)<br />
Contains a <a href="#space_manager_device">Space manager device</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;" colspan="4"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>144</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Flags<br />
See section: <a href="#space_manager_flags">Space manager flags</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>148</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (sm_ip_bm_tx_multiplier)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>152</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (sm_ip_block_count)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>160</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (sm_ip_bm_size_in_blocks)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>164</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (sm_ip_bm_block_count)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>168</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (sm_ip_bm_base)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>176</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (sm_ip_base)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>184</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (sm_fs_reserve_block_count)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>192</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (sm_fs_reserve_alloc_count)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Space manager free queues (sm_fq)</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>200</p></td>
<td style="text-align: left;"><p>40</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown space free queue (SFQ_IP)</span></strong><br />
Contains a <a href="#space_manager_free_queue">Space manager free queue</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>240</p></td>
<td style="text-align: left;"><p>40</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Main space free queue (SFQ_MAIN)*<br />
Contains a <a href="#space_manager_free_queue">Space manager free queue</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>280</p></td>
<td style="text-align: left;"><p>40</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Tier2 space free queue (SFQ_TIER2)*<br />
Contains a <a href="#space_manager_free_queue">Space manager free queue</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;" colspan="4"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>320</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (sm_ip_bm_free_head)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>322</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (sm_ip_bm_free_tail)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>324</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (sm_ip_bm_xid_offset)</span></strong><br />
Contains an offset in bytes relative to the start of the space manager</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>328</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (sm_ip_bitmap_offset)</span></strong><br />
Contains an offset in bytes relative to the start of the space manager</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>332</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (sm_ip_bm_free_next_offset)</span></strong><br />
Contains an offset in bytes relative to the start of the space manager</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>336</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>1</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (sm_version)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>340</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (sm_struct_size)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Space manager data zone (sm_datazone)</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>344</p></td>
<td style="text-align: left;"><p>8 x 72</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Main allocation zones</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>920</p></td>
<td style="text-align: left;"><p>8 x 72</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Tier2 allocation zones</p></td>
</tr>
<tr class="even">
<td style="text-align: left;" colspan="4"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>1492</p></td>
<td style="text-align: left;"><p>…</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (data)</span></strong></p></td>
</tr>
</tbody>
</table>

### <span id="space_manager_flags"></span>Space manager flags

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0x00000001</p></td>
<td style="text-align: left;"><p>SM_FLAG_VERSIONED</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
</tbody>
</table>

### <span id="space_manager_device"></span>Space manager device

A space manager device (spaceman\_device\_t) is 48 bytes of size and
consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of blocks (sm_block_count)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of chunks (sm_chunk_count)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of chunk information blocks (CIBs) (sm_cib_count)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>20</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of chunk information address blocks (CABs) (sm_cab_count)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of unused blocks (sm_free_count)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>32</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (sm_addr_offset)</span></strong><br />
Contains an offset in bytes relative to the start of the space manager</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>36</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (sm_reserved)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>40</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (sm_reserved2)</span></strong></p></td>
</tr>
</tbody>
</table>

### <span id="space_manager_free_queue"></span>Space manager free queue

A space manager free queue (spaceman\_free\_queue\_t) is 40 bytes of
size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (sfq_count)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Space manager free queue tree object identifier (sfq_tree_oid)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Space manager free queue oldest transaction identifier (sfq_oldest_xid)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (sfq_tree_node_limit)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>26</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (sfq_pad16)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>28</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (sfq_pad32)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>32</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (sfq_reserved)</span></strong></p></td>
</tr>
</tbody>
</table>

### <span id="space_manager_allocation_zone"></span>Space manager allocation zone

A space manager allocation zone
(spaceman\_allocation\_zone\_info\_phys\_t) is 72 bytes of size and
consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Current allocation zone boundaries (saz_current_boundaries)<br />
Contains <a href="#space_manager_zone_boundaries">Space manager zone boundaries</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>7 x 8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Previous allocation zone boundaries (saz_previous_boundaries)<br />
Contains <a href="#space_manager_zone_boundaries">Space manager zone boundaries</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>64</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (saz_zone_id)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>66</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (saz_previous_boundary_index)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>68</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (saz_reserved)</span></strong></p></td>
</tr>
</tbody>
</table>

### <span id="space_manager_zone_boundaries"></span>Space manager zone\_boundaries

A space manager zone boundaries
(spaceman\_allocation\_zone\_boundaries\_t) is 8 bytes of size and
consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (saz_zone_start)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (saz_zone_end)</span></strong></p></td>
</tr>
</tbody>
</table>

### Notes

sm\_addr\_offset points to block number which points to a
OBJECT\_TYPE\_SPACEMAN\_CIB block. Probably an
OBJECT\_TYPE\_SPACEMAN\_CAB block when necessary.

    00052000  0d cd df 3f cb 2a 20 80  4d 00 00 00 00 00 00 00  |...?.* .M.......|
    00052010  04 00 00 00 00 00 00 00  07 00 00 40 00 00 00 00  |...........@....|
    00052020  00 00 00 00 01 00 00 00  04 00 00 00 00 00 00 00  |................|
    00052030  00 00 00 00 00 00 00 00  f6 03 00 00 86 03 00 00  |................|
    00052040  4e 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |N...............|
    00052050  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
    *
    00053000  ff ff ff ff ff ff ff ff  ff ff ff ff ff ff 00 00  |................|
    00053010  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
    *

## <span id="chunk_information_address_block"></span>Chunk information address block

The chunk information address block (cib\_addr\_block\_t) is variable of
size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object header (cab_o)</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object checksum<br />
See section: <a href="#object_checksum">Object checkum</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object identifier</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object transaction identifier (xid)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x40000006</p></td>
<td style="text-align: left;"><p>Object type<br />
See section: <a href="#object_types">Object types</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>28</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x00000000</p></td>
<td style="text-align: left;"><p>Object subtype</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object values</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>32</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (cab_index)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>36</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of chunk information blocks (CIBs) (cab_cib_count)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;" colspan="4"><p><em>Chunk information block physical addresses (cab_cib_addr)</em></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>40</p></td>
<td style="text-align: left;"><p>8 x Number of CIBs</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Physical address of chunk information blocks (CIB)</p></td>
</tr>
</tbody>
</table>

## <span id="chunk_information_block"></span>Chunk information block

The chunk information block (chunk\_info\_block\_t) is variable of size
and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object header (cib_o)</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object checksum<br />
See section: <a href="#object_checksum">Object checkum</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object identifier</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object transaction identifier (xid)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x40000007</p></td>
<td style="text-align: left;"><p>Object type<br />
See section: <a href="#object_types">Object types</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>28</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x00000000</p></td>
<td style="text-align: left;"><p>Object subtype</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object values</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>32</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (cib_index)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>36</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of chunk information entries (cib_chunk_info_count)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;" colspan="4"><p><em>Chunk information entries (cib_chunk_info)</em></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>40</p></td>
<td style="text-align: left;"><p>8 x Number of entries</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Array of chunk information entries</p></td>
</tr>
</tbody>
</table>

### <span id="chunk_information_entry"></span>Chunk information entry

The chunk information entry (chunk\_info\_t) is 32 bytes of size and
consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (ci_xid)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (ci_addr)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (ci_block_count)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>20</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (ci_free_count)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (ci_bitmap_addr)</span></strong></p></td>
</tr>
</tbody>
</table>

## <span id="reaper"></span>Reaper

The reaper is **<span class="yellow-background">unknown</span>** of size
and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object header</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object checksum<br />
See section: <a href="#object_checksum">Object checkum</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object identifier</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object transaction identifier (xid)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x80000011</p></td>
<td style="text-align: left;"><p>Object type<br />
See section: <a href="#object_types">Object types</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>28</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x00000000</p></td>
<td style="text-align: left;"><p>Object subtype</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object values</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>32</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
</tbody>
</table>

### <span id="reaper_list"></span>Reaper list

The reaper list entry is **<span
class="yellow-background">unknown</span>** of size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object header</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object checksum<br />
See section: <a href="#object_checksum">Object checkum</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object identifier</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object transaction identifier (xid)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x80000012</p></td>
<td style="text-align: left;"><p>Object type<br />
See section: <a href="#object_types">Object types</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>28</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x00000000</p></td>
<td style="text-align: left;"><p>Object subtype</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object values</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>32</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>36</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>40</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>44</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (max_record_count)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>48</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (record_count)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>52</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (first_index)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>56</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (last_index)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>60</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (free_index)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>64</p></td>
<td style="text-align: left;"><p>100 x …</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Array of reaper list entries (nrle)<br />
See section: <a href="#reaper_list_entry">Reaper list entry</a></p></td>
</tr>
</tbody>
</table>

### <span id="reaper_list_entry"></span>Reaper list entry

The reaper list entry is 40 bytes of size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Forward link (fwlink)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Type (type)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>12</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Block size (blksize)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object identifier (oid)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Physical address (paddr)<br />
Contains a block number relative to the start of the container</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>32</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object transaction identifier (xid)</p></td>
</tr>
</tbody>
</table>

## <span id="key_bag"></span>Key bag

The key bag consists of:

-   Container or volume key bag object

-   Key bag header

-   Key bag entries

## Container key bag object

The container key bag object contains key data of the container.

The container key bag object is 32 bytes of size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object header</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object checksum<br />
See section: <a href="#object_checksum">Object checkum</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object identifier</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object transaction identifier (xid)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x6b657973 ("syek")</p></td>
<td style="text-align: left;"><p>Object type<br />
See section: <a href="#object_types">Object types</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>28</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x00000000</p></td>
<td style="text-align: left;"><p>Object subtype</p></td>
</tr>
</tbody>
</table>

## Volume key bag object

The volume key bag object contains key data of a specific volume.

The volume key bag object is 32 bytes of size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object header</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object checksum<br />
See section: <a href="#object_checksum">Object checkum</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object identifier</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object transaction identifier (xid)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x72656373 ("scer")</p></td>
<td style="text-align: left;"><p>Object type<br />
See section: <a href="#object_types">Object types</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>28</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x00000000</p></td>
<td style="text-align: left;"><p>Object subtype</p></td>
</tr>
</tbody>
</table>

## Key bag header

The key bag header (kb\_locker\_t) is 16 bytes of size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"><p>Format version (kl_version)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of entries (kl_nkeys)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Key bag data size (kl_nbytes)<br />
Contains the size of the key bag data, this includes the size of key bag header</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (padding)</span></strong></p></td>
</tr>
</tbody>
</table>

## <span id="key_bag_entries"></span>Key bag entries

A key bag entry consists of:

-   a key bag entry header

-   a key bag entry data

-   alignment padding

The key bag entry header specifies the type of the key bag entry data.

The key bag entries are 16-byte aligned.

### Key bag entry header

The key bag entry header (keybag\_entry\_t) is 24 bytes of size and
consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Volume identifer (ke_uuid)<br />
Contains a UUID stored in big-endian</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Entry type (ke_tag)<br />
See section: <a href="#key_bag_entry_types">Key bag entry types</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>18</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Entry data size (ke_keylen)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>20</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (padding)</span></strong></p></td>
</tr>
</tbody>
</table>

### <span id="key_bag_entry_types"></span>Key bag entry types

#### Container key bag entry types

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0x00</p></td>
<td style="text-align: left;"><p>KB_TAG_UNKNOWN</p></td>
<td style="text-align: left;"><p>Unknown</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x01</p></td>
<td style="text-align: left;"><p>KB_TAG_WRAPPING_KEY</p></td>
<td style="text-align: left;"><p>Wrapping key</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x02</p></td>
<td style="text-align: left;"><p>KB_TAG_VOLUME_KEY</p></td>
<td style="text-align: left;"><p>Volume master key<br />
See section: <a href="#key_bag_kek_packed_object">Key encrypted key (KEK) packed object</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x03</p></td>
<td style="text-align: left;"><p>KB_TAG_VOLUME_UNLOCK_RECORDS</p></td>
<td style="text-align: left;"><p>Volume key bag extent<br />
See section: <a href="#key_bag_data_extent">Key bag data extent</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x04</p></td>
<td style="text-align: left;"><p>KB_TAG_VOLUME_PASSPHRASE_HINT</p></td>
<td style="text-align: left;"><p>Passphrase hint</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0xf8</p></td>
<td style="text-align: left;"><p>KB_TAG_USER_PAYLOAD</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (user payload)</span></strong></p></td>
</tr>
</tbody>
</table>

The volume master key is encryped with a volume key.

#### Volume key bag entry types

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>3</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Volume key<br />
See section: <a href="#key_bag_kek_packed_object">Key encrypted key (KEK) packed object</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Password Hint<br />
Contains a string without end-of-string character</p></td>
</tr>
</tbody>
</table>

The volume key is encryped with an user key.

### <span id="key_bag_packed_object"></span>Key bag packed object

The packed object consist of an object packed value that embeds
attribute packed values.

#### Key bag packed value

The key bag packed value is variable of size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>1</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Value tag (or value type)<br />
<strong><span class="yellow-background">Unknown (Where the most-significant bit represents a user-defined flag?)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>1</p></td>
<td style="text-align: left;"><p>1</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Value data size<br />
If the most-significant bit is set the value data size is stored in the next ( value &amp; 0x7f ) bytes<br />
Seen: 0x81</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>…</p></td>
<td style="text-align: left;"><p>…</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Value data</p></td>
</tr>
</tbody>
</table>

The meaning of the value tags differ per packed object type.

A packed value with a tag and size of 0 signifies the end of the packed
values.

#### <span id="key_bag_kek_packed_object"></span>Key encrypted key (KEK) packed object

The packed object value tag of a key encrypted key is 0x30 and contains
the following attribute value tags:

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0x80</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x81</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>HMAC</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x82</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (salt?)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0xa3</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Wrapped Wrapped Key Encrypted Key (KEK) packed object<br />
See section: <a href="#key_bag_wrapped_kek_packed_object">Wrapped Key Encrypted Key (KEK) packed object</a></p></td>
</tr>
</tbody>
</table>

#### <span id="key_bag_wrapped_kek_packed_object"></span>Wrapped Key Encrypted Key (KEK) packed object

The packed object value tag of a wrapped kek encrypted key is 0xa3 and
contains the following attribute value tags:

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0x80</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x81</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Volume identifer<br />
Contains a UUID stored in big-endian</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x82</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Wrapped Key Encrypted Key (KEK) metadata<br />
See section: <a href="#wrapped_kek_metadata">Wrapped Key Encrypted Key (KEK) metadata</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x83</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Wrapped Key Encrypted Key (KEK) data</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x84</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of iterations</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x85</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Salt for the PBKDF2 algorithm</p></td>
</tr>
</tbody>
</table>

### <span id="wrapped_kek_metadata"></span>Wrapped Key Encrypted Key (KEK) metadata

The Wrapped Key Encrypted Key (KEK) metadata is 8 bytes of size and
consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Encryption method<br />
See section: <a href="#encryption_methods">Encryption methods</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>6</p></td>
<td style="text-align: left;"><p>1</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>7</p></td>
<td style="text-align: left;"><p>1</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
</tbody>
</table>

#### <span id="encryption_methods"></span>Encryption methods

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (AES-256)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (AES-128 FVDE (CoreStorage FileVault) compatible)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (AES-256)</span></strong><br />
Seen in combination with recovery password protected volume key</p></td>
</tr>
</tbody>
</table>

### <span id="key_bag_data_extent"></span>Key bag data extent

The key bag data extent is 16 bytes of size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Key bag block number</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Key bag number of blocks</p></td>
</tr>
</tbody>
</table>

## Volume

The volume consists of:

-   volume superblock

-   volume object map

-   …

Individual APFS volume have a corresponding "synthesized" device file
though this cannot be directly read.

## <span id="volume_superblock"></span>Volume superblock

The volume superblock (apfs\_superblock\_t) is 940 bytes of size and
consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object header</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object checksum<br />
See section: <a href="#object_checksum">Object checkum</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object identifier</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object transaction identifier (xid)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x0000000d<br />
0x4000000d (for snapshots)</p></td>
<td style="text-align: left;"><p>Object type<br />
See section: <a href="#object_types">Object types</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>28</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x00000000</p></td>
<td style="text-align: left;"><p>Object subtype</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object values</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>32</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>"APSB"</p></td>
<td style="text-align: left;"><p>Signature (apfs_magic)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>36</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (apfs_fs_index)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>40</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Compatible feature flags (apfs_features)<br />
<a href="#volume_superblock_feature">Volume superblock features</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>48</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Read-only compatible feature flags (apfs_readonly_compatible_features)<br />
<a href="#volume_read_only_superblock_feature">Volume read-only superblock features</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>56</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Incompatible feature flags (apfs_incompatible_features)<br />
<a href="#volume_incompatible_superblock_feature">Volume incompatible superblock features</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>64</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (apfs_unmount_time)</span></strong><br />
Signed integer that contains the number of nanoseconds since January 1, 1970 00:00:00 UTC or 0 if not set</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>72</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of reserved blocks (apfs_reserve_block_count)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>80</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of quota blocks (apfs_quota_block_count)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>88</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (apfs_fs_alloc_count)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>96</p></td>
<td style="text-align: left;"><p>20</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Encryption state (apfs_meta_crypto)<br />
See section: <a href="#encryption_state">Encryption state</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>116</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>File system root tree object type (apfs_root_tree_type)<br />
See section: <a href="#object_types">Object types</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>120</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Extent-reference tree object type (apfs_extentref_tree_type)<br />
See section: <a href="#object_types">Object types</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>124</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Snapshot metadata tree object type (apfs_snap_meta_tree_type)<br />
See section: <a href="#object_types">Object types</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>132</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object map block number (apfs_omap_oid)<br />
Contains a block number relative to the start of the container of the <a href="#object_map">object map</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>140</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>File system root tree object identifier (apfs_root_tree_oid)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>148</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Extent-reference tree block number (apfs_extentref_tree_oid)<br />
See section: <a href="#extent_reference_tree">Extent reference tree</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>156</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Snapshot metadata tree block number (apfs_snap_meta_tree_oid)<br />
See section: <a href="#snapshot_metadata_tree">Snapshot metadata tree</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>164</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (apfs_revert_to_xid)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>172</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (apfs_revert_to_sblock_oid)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>180</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Next (available) file system object identifier (apfs_next_obj_id)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>188</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (apfs_num_files)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>196</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (apfs_num_directories)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>204</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (apfs_num_symlinks)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>212</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (apfs_num_other_fsobjects)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>220</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (apfs_num_snapshots)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>228</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (apfs_total_blocks_alloced)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>236</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (apfs_total_blocks_freed)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>244</p></td>
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Volume identifier (apfs_vol_uuid)<br />
Contains a UUID stored in big-endian</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>260</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Modification date and time (apfs_last_mod_time)<br />
Signed integer that contains the number of nanoseconds since January 1, 1970 00:00:00 UTC or 0 if not set</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>268</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Volume flags (apfs_fs_flags)<br />
See section: <a href="#volume_superblock_flags">Volume superblock flags</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>276</p></td>
<td style="text-align: left;"><p>40</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Creation change information (apfs_formatted_by)<br />
See section: <a href="#change_information">Change information</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>316</p></td>
<td style="text-align: left;"><p>8 x 40</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Modification change information (apfs_modified_by)<br />
<strong><span class="yellow-background">Contains the 8 last entries from least recent to most recent?</span></strong><br />
See section: <a href="#change_information">Change information</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>636</p></td>
<td style="text-align: left;"><p>256</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Volume name (apfs_volname)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>892</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Next (available) document identifier (apfs_next_doc_id)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>896</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (apfs_role)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>898</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (reserved)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>900</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (apfs_root_to_xid)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>908</p></td>
<td style="text-align: left;"><p>32</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (apfs_er_state_oid)</span></strong></p></td>
</tr>
</tbody>
</table>

## <span id="encryption_state"></span>Encryption state

The encryption state (wrapped\_meta\_crypto\_state\_t) is 20 bytes of
size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Major format version (major_version)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Minor format version (minor_version)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Flags (cpflags)<br />
See section: <a href="#encryption_state_flags">Encryption state flags</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (persistent_class)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>12</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (key_os_version)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (key_revision)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>18</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (unused)</span></strong></p></td>
</tr>
</tbody>
</table>

### <span id="encryption_state_flags"></span>Encryption state flags

**<span class="yellow-background">TODO: complete this section.</span>**

## <span id="change_information"></span>Change information

The change information (apfs\_modified\_by\_t) is 48 bytes of size and
consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>32</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Application (id)<br />
String that contains the first 31 characters of the name and version of the application that changed the file system<br />
Contains 0 if not set</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>32</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Change date and time (timestamp)<br />
Signed integer that contains the number of nanoseconds since January 1, 1970 00:00:00 UTC or 0 if not set</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>40</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Change object transaction number (last_xid)<br />
Contains 0 if not set</p></td>
</tr>
</tbody>
</table>

### <span id="volume_superblock_flags"></span>Volume flags

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0x0000000000000001</p></td>
<td style="text-align: left;"><p>APFS_FS_UNENCRYPTED</p></td>
<td style="text-align: left;"><p>Volume is unencrypted</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0000000000000002</p></td>
<td style="text-align: left;"><p>APFS_FS_EFFACEABLE</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (Volume supports effaceable storage?)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x0000000000000004</p></td>
<td style="text-align: left;"><p>APFS_FS_RESERVED_4</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (reserved)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0000000000000008</p></td>
<td style="text-align: left;"><p>APFS_FS_ONEKEY</p></td>
<td style="text-align: left;"><p>Volume uses software encryption with a single key (volume master key)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x0000000000000010</p></td>
<td style="text-align: left;"><p>APFS_FS_SPILLEDOVER</p></td>
<td style="text-align: left;"><p>Volume has run out of allocated space on the solid-state drive</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0000000000000020</p></td>
<td style="text-align: left;"><p>APFS_FS_RUN_SPILLOVER_CLEANER</p></td>
<td style="text-align: left;"><p>Volume has spilled over and the spillover cleaner must be run</p></td>
</tr>
</tbody>
</table>

### <span id="volume_superblock_feature"></span>Volume superblock features

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0x0000000000000001</p></td>
<td style="text-align: left;"><p>APFS_FEATURE_DEFRAG_PRERELEASE</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0000000000000002</p></td>
<td style="text-align: left;"><p>APFS_FEATURE_HARDLINK_MAP_RECORDS</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x0000000000000004</p></td>
<td style="text-align: left;"><p>APFS_FEATURE_DEFRAG</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
</tbody>
</table>

### <span id="volume_read_only_feature_flags"></span>Volume read-only feature flags

Current no read-only feature flags are defined

### <span id="volume_incompatible_feature_flags"></span>Volume incompatible feature flags

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0x0000000000000001</p></td>
<td style="text-align: left;"><p>APFS_INCOMPAT_CASE_INSENSITIVE</p></td>
<td style="text-align: left;"><p>Filenames are case insensitive</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0000000000000002</p></td>
<td style="text-align: left;"><p>APFS_INCOMPAT_DATALESS_SNAPS</p></td>
<td style="text-align: left;"><p>Volume contains one or more snapshots without data</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x0000000000000004</p></td>
<td style="text-align: left;"><p>APFS_INCOMPAT_ENC_ROLLED</p></td>
<td style="text-align: left;"><p>Encryption keys of the volume have been changed</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0000000000000008</p></td>
<td style="text-align: left;"><p>APFS_INCOMPAT_NORMALIZATION_INSENSITIVE</p></td>
<td style="text-align: left;"><p>Filenames are normalization insensitive</p></td>
</tr>
</tbody>
</table>

## <span id="file_system"></span>File system

The file system structures are stored in a [B-tree](#btree).

The file system B-tree uses identifiers similar to catalog identifiers
(CNIDs) on HFS/HFS+/HFSX. In this document these identifiers are
referred to as File System object identifiers (FSOIDs) to contrast other
object identifiers (OIDs).

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">FSOID</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Assignment</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (Reserved)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>1</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Parent identifier of the root directory (folder), nameless</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Directory identifier of the root directory (folder), named "root"</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>3</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong>, named "private-dir"</p></td>
</tr>
</tbody>
</table>

## File system B-tree key

The file system B-tree key is variable of size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object identifier and type (obj_id_and_type)</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>60 bits</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>File system object identifier (FSOID)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>7.4</p></td>
<td style="text-align: left;"><p>4 bits</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>File system data type<br />
See section: <a href="#file_system_data_types">File system data types</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>…</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Optional additional key data dependent on the data type</p></td>
</tr>
</tbody>
</table>

## <span id="file_system_data_types"></span>File system data types

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0x0</p></td>
<td style="text-align: left;"><p>APFS_TYPE_ANY</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (Any)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x1</p></td>
<td style="text-align: left;"><p>APFS_TYPE_SNAP_METADATA</p></td>
<td style="text-align: left;"><p>Snapshot metadata See section: <a href="#snapshot_metadata">Snapshot metadata</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x2</p></td>
<td style="text-align: left;"><p>APFS_TYPE_EXTENT</p></td>
<td style="text-align: left;"><p>Extent<br />
See section: <a href="#extent">Extent</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x3</p></td>
<td style="text-align: left;"><p>APFS_TYPE_INODE</p></td>
<td style="text-align: left;"><p>Inode<br />
See section: <a href="#inode">Inode</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x4</p></td>
<td style="text-align: left;"><p>APFS_TYPE_XATTR</p></td>
<td style="text-align: left;"><p>Extended attribute (xattr)<br />
See section: <a href="#extended_attribute">Extended attribute</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x5</p></td>
<td style="text-align: left;"><p>APFS_TYPE_SIBLING_LINK</p></td>
<td style="text-align: left;"><p>Sibling link<br />
See section: <a href="#sibling_link">Sibling link</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x6</p></td>
<td style="text-align: left;"><p>APFS_TYPE_DSTREAM_ID</p></td>
<td style="text-align: left;"><p>Data stream identifier<br />
See section: <a href="#data_stream_identifier">Data stream identifier</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x7</p></td>
<td style="text-align: left;"><p>APFS_TYPE_CRYPTO_STATE</p></td>
<td style="text-align: left;"><p>Encryption state See section: <a href="#encryption_state">Encryption state</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x8</p></td>
<td style="text-align: left;"><p>APFS_TYPE_FILE_EXTENT</p></td>
<td style="text-align: left;"><p>File extent See section: <a href="#file_extent">File extent</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x9</p></td>
<td style="text-align: left;"><p>APFS_TYPE_DIR_REC</p></td>
<td style="text-align: left;"><p>Directory record<br />
See section: <a href="#directory_record">Directory record</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0xa</p></td>
<td style="text-align: left;"><p>APFS_TYPE_DIR_STATS</p></td>
<td style="text-align: left;"><p>Directory stats<br />
See section: <a href="#directory_stats">Directory stats</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0xb</p></td>
<td style="text-align: left;"><p>APFS_TYPE_SNAP_NAME</p></td>
<td style="text-align: left;"><p>Snapshot name<br />
See section: <a href="#snapshot_name">Snapshot name</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0xc</p></td>
<td style="text-align: left;"><p>APFS_TYPE_SIBLING_MAP</p></td>
<td style="text-align: left;"><p>Sibling map<br />
See section: <a href="#sibling_map">Sibling map</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0xf</p></td>
<td style="text-align: left;"><p>APFS_TYPE_INVALID</p></td>
<td style="text-align: left;"><p>Invalid</p></td>
</tr>
</tbody>
</table>

## File system B-tree branch node value

A file system B-tree node contains branch node values if BTNODE\_LEAF is
not set. The corresponding file system B-tree key represents the first
key in the branch.

A file system B-tree branch node value is 8 bytes of size and consists
of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>B-tree sub node object identifier<br />
The object identifiers can be resolved in the <a href="#object_map">object map</a> to a "physical" location</p></td>
</tr>
</tbody>
</table>

## <span id="snapshot_metadata"></span>Snapshot metadata

The snapshot metadata value (j\_snap\_metadata\_val\_t) is variable of
size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Extent-reference tree block number<br />
Contains a block number relative to the start of the container</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Volume superblock block number<br />
Contains a block number relative to the start of the container</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Creation time<br />
Signed integer that contains the number of nanoseconds since January 1, 1970 00:00:00 UTC or 0 if not set</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Change (or last modification) time<br />
Signed integer that contains the number of nanoseconds since January 1, 1970 00:00:00 UTC or 0 if not set</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>32</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (inum)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>40</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Extent-reference tree object type (extentref_tree_type)<br />
See section: <a href="#object_types">Object types</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>44</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Flags<br />
See section: <a href="#snapshot_metadata_flags">Snapshot metadata flags</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>48</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Name string size (name_len)<br />
Includes the size of the end-of-string character</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>50</p></td>
<td style="text-align: left;"><p>…</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Name string (name)<br />
Contains an UTF-8 encoded string with an end-of-string character</p></td>
</tr>
</tbody>
</table>

### <span id="snapshot_metadata_flags"></span>Snapshot metadata flags

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0x00000001</p></td>
<td style="text-align: left;"><p>SNAP_META_PENDING_DATALESS</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
</tbody>
</table>

## <span id="extent"></span>Extent

### Extent key data

The extent key data (j\_phys\_ext\_key\_t) is 8 bytes of size and
consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>60 bits</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>File system object identifier (FSOID)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>7.4</p></td>
<td style="text-align: left;"><p>4 bits</p></td>
<td style="text-align: left;"><p>0x2</p></td>
<td style="text-align: left;"><p>File system data type<br />
See section: <a href="#file_system_data_types">File system data types</a></p></td>
</tr>
</tbody>
</table>

### Extent value data

The extent value data (j\_phys\_ext\_val\_t) is 20 bytes of size and
consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Extent size and data type (len_and_kind)</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>60 bits</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Extent data size</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>7.4</p></td>
<td style="text-align: left;"><p>4 bits</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>File system data type<br />
See section: <a href="#file_system_data_types">File system data types</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;" colspan="4"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>File system object identifier of owner (owning_obj_id)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Reference count (refcnt)</p></td>
</tr>
</tbody>
</table>

## <span id="inode"></span>Inode

### Inode key data

The inode key data (j\_inode\_key\_t) is 8 bytes of size and consists
of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>60 bits</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>File system object identifier (FSOID)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>7.4</p></td>
<td style="text-align: left;"><p>4 bits</p></td>
<td style="text-align: left;"><p>0x3</p></td>
<td style="text-align: left;"><p>File system data type<br />
See section: <a href="#file_system_data_types">File system data types</a></p></td>
</tr>
</tbody>
</table>

### Inode value data

The inode value data (APFS\_TYPE\_INVALID) is 8 bytes of size and
consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Parent file system object identifier (parent_id)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Data stream file system object identifier (private_id)<br />
Contains the file system object identifier of the file extents that make up the data stream</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Creation date and time (create_time)<br />
Signed integer that contains the number of nanoseconds since January 1, 1970 00:00:00 UTC or 0 if not set</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Modification date and time (mod_time)<br />
Signed integer that contains the number of nanoseconds since January 1, 1970 00:00:00 UTC or 0 if not set</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>32</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Inode change date and time (change_time)<br />
Signed integer that contains the number of nanoseconds since January 1, 1970 00:00:00 UTC or 0 if not set</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>48</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Access date and time (access_time)<br />
Signed integer that contains the number of nanoseconds since January 1, 1970 00:00:00 UTC or 0 if not set</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>56</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Inode flags (internal_flags)<br />
See section: <a href="#inode_flags">Inode flags</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>64</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of children (nchildren) or number of links (nlink)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>68</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (default_protection_class)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>72</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (write_generation_counter)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>76</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>BSD file entry flags (bsd_flags)<br />
See section: <a href="#bsd_file_entry_flags">BSD file entry flags</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>80</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Owner user identifier (owner)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>84</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Group identifier (gid)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>86</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>File mode<br />
See section: <a href="#file_modes">File modes</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>88</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (pad1)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>90</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (pad2)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>98</p></td>
<td style="text-align: left;"><p>…</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Extended fields (xfields)<br />
See section: <a href="#extended_fields">Extended fields</a></p></td>
</tr>
</tbody>
</table>

#### <span id="inode_flags"></span>Inode flags

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0x0000000000000001</p></td>
<td style="text-align: left;"><p>INODE_IS_APFS_PRIVATE</p></td>
<td style="text-align: left;"><p>Is private<br />
The inode is used internally, typically for a data stream</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0000000000000002</p></td>
<td style="text-align: left;"><p>INODE_MAINTAIN_DIR_STATS</p></td>
<td style="text-align: left;"><p>Maintains directory stats<br />
The inode tracks the size of all of its children</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x0000000000000004</p></td>
<td style="text-align: left;"><p>INODE_DIR_STATS_ORIGIN</p></td>
<td style="text-align: left;"><p>Maintains directory stats explicitly set, not inherited<br />
The inode has the INODE_MAINTAIN_DIR_STATS flag set explicitly, not due to inheritance</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0000000000000008</p></td>
<td style="text-align: left;"><p>INODE_PROT_CLASS_EXPLICIT</p></td>
<td style="text-align: left;"><p>Protection class explicitly set, not inherited<br />
The inode data protection class was set explicitly when the inode was created</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x0000000000000010</p></td>
<td style="text-align: left;"><p>INODE_WAS_CLONED</p></td>
<td style="text-align: left;"><p>Was cloned<br />
The inode was created by cloning another inode</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0000000000000020</p></td>
<td style="text-align: left;"><p>INODE_FLAG_UNUSED</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (Reserved)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x0000000000000040</p></td>
<td style="text-align: left;"><p>INODE_HAS_SECURITY_EA</p></td>
<td style="text-align: left;"><p>Has security extended attribute<br />
The inode has an access control list</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0000000000000080</p></td>
<td style="text-align: left;"><p>INODE_BEING_TRUNCATED</p></td>
<td style="text-align: left;"><p>Is truncated<br />
The inode was truncated</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x0000000000000100</p></td>
<td style="text-align: left;"><p>INODE_HAS_FINDER_INFO</p></td>
<td style="text-align: left;"><p>Has Finder information<br />
The inode has a Finder info extended field</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0000000000000200</p></td>
<td style="text-align: left;"><p>INODE_IS_SPARSE</p></td>
<td style="text-align: left;"><p>Is sparse<br />
The inode has a sparse byte count extended field</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x0000000000000400</p></td>
<td style="text-align: left;"><p>INODE_WAS_EVER_CLONED</p></td>
<td style="text-align: left;"><p>Was cloned<br />
The inode has been cloned at least once</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0000000000000800</p></td>
<td style="text-align: left;"><p>INODE_ACTIVE_FILE_TRIMMED</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (TODO)</span></strong><br />
The inode is an overprovisioning file that has been trimmed</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x0000000000001000</p></td>
<td style="text-align: left;"><p>INODE_PINNED_TO_MAIN</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (TODO)</span></strong><br />
The inode file content is always on the main storage device<br />
This flag is used for Fusion drives where the main storage is a solid-state drive</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0000000000002000</p></td>
<td style="text-align: left;"><p>INODE_PINNED_TO_TIER2</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (TODO)</span></strong><br />
The inode file content is always on the secondary storage device<br />
This flag is used for Fusion drives where the secondary storage is a (magnetic) hard drive</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x0000000000004000</p></td>
<td style="text-align: left;"><p>INODE_HAS_RSRC_FORK</p></td>
<td style="text-align: left;"><p>Has resource fork<br />
The inode has a resource fork</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0000000000008000</p></td>
<td style="text-align: left;"><p>INODE_NO_RSRC_FORK</p></td>
<td style="text-align: left;"><p>Has no resource fork<br />
The inode does not have a resource fork</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x0000000000010000</p></td>
<td style="text-align: left;"><p>INODE_ALLOCATION_SPILLEDOVER</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (TODO)</span></strong><br />
The inode file content has some space allocated outside of the preferred storage tier for that file</p></td>
</tr>
</tbody>
</table>

#### <span id="file_modes"></span>File modes

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0xf000 (0170000)</p></td>
<td style="text-align: left;"><p>S_IFMT</p></td>
<td style="text-align: left;"><p>File type bitmask</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x1000 (0010000)</p></td>
<td style="text-align: left;"><p>S_IFIFO</p></td>
<td style="text-align: left;"><p>Named pipe</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x2000 (0020000)</p></td>
<td style="text-align: left;"><p>S_IFCHR</p></td>
<td style="text-align: left;"><p>Character-special file (Character device)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x4000 (0040000)</p></td>
<td style="text-align: left;"><p>S_IFDIR</p></td>
<td style="text-align: left;"><p>Directory</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x6000 (0060000)</p></td>
<td style="text-align: left;"><p>S_IFBLK</p></td>
<td style="text-align: left;"><p>Block-special file (Block device)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x8000 (0100000)</p></td>
<td style="text-align: left;"><p>S_IFREG</p></td>
<td style="text-align: left;"><p>Regular file</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0xa000 (0120000)</p></td>
<td style="text-align: left;"><p>S_IFLNK</p></td>
<td style="text-align: left;"><p>Symbolic link</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0xc000 (0140000)</p></td>
<td style="text-align: left;"><p>S_IFSOCK</p></td>
<td style="text-align: left;"><p>Socket</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0xe000 (0160000)</p></td>
<td style="text-align: left;"><p>S_IFWHT</p></td>
<td style="text-align: left;"><p>Whiteout<br />
A whiteout is a file entry that covers up all entries of a particular name from lower branches</p></td>
</tr>
</tbody>
</table>

#### <span id="bsd_file_entry_flags"></span>BSD file entry flags

The BSD file entry flags are defined in the *&lt;sys/stat.h>* header
file.

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0x0000ffff</p></td>
<td style="text-align: left;"><p>UF_SETTABLE</p></td>
<td style="text-align: left;"><p>bitmask of owner changeable flags</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00000001</p></td>
<td style="text-align: left;"><p>UF_NODUMP</p></td>
<td style="text-align: left;"><p>do not dump file entry</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00000002</p></td>
<td style="text-align: left;"><p>UF_IMMUTABLE</p></td>
<td style="text-align: left;"><p>file entry is immutable and may not be changed</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00000004</p></td>
<td style="text-align: left;"><p>UF_APPEND</p></td>
<td style="text-align: left;"><p>writes to file entry may only append</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00000008</p></td>
<td style="text-align: left;"><p>UF_OPAQUE</p></td>
<td style="text-align: left;"><p>directory is opaque wrt. union</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00000010</p></td>
<td style="text-align: left;"><p>UF_NOUNLINK</p></td>
<td style="text-align: left;"><p>file entry may not be removed or renamed<br />
Not implement in MacOS</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00000020</p></td>
<td style="text-align: left;"><p>UF_COMPRESSED</p></td>
<td style="text-align: left;"><p>file entry is compressed</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00000040</p></td>
<td style="text-align: left;"><p>UF_TRACKED</p></td>
<td style="text-align: left;"><p>notify about file entry changes</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00000080</p></td>
<td style="text-align: left;"><p>UF_DATAVAULT</p></td>
<td style="text-align: left;"><p>entitlement required for reading and writing</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00008000</p></td>
<td style="text-align: left;"><p>UF_HIDDEN</p></td>
<td style="text-align: left;"><p>file entry is hidden</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0xffff0000</p></td>
<td style="text-align: left;"><p>SF_SETTABLE</p></td>
<td style="text-align: left;"><p>bitmask of superuser changeable flags</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x001f0000</p></td>
<td style="text-align: left;"><p>SF_SUPPORTED</p></td>
<td style="text-align: left;"><p>bitmask of superuser supported flags</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00010000</p></td>
<td style="text-align: left;"><p>SF_ARCHIVED</p></td>
<td style="text-align: left;"><p>file entry is archived</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00020000</p></td>
<td style="text-align: left;"><p>SF_IMMUTABLE</p></td>
<td style="text-align: left;"><p>file entry is immutable and may not be changed</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00040000</p></td>
<td style="text-align: left;"><p>SF_APPEND</p></td>
<td style="text-align: left;"><p>writes to file entry may only append</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00080000</p></td>
<td style="text-align: left;"><p>SF_RESTRICTED</p></td>
<td style="text-align: left;"><p>entitlement required for writing</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x00100000</p></td>
<td style="text-align: left;"><p>SF_NOUNLINK</p></td>
<td style="text-align: left;"><p>file entry may not be removed, renamed or used as mount point</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x00200000</p></td>
<td style="text-align: left;"><p>SF_SNAPSHOT</p></td>
<td style="text-align: left;"><p>snapshot inode<br />
Not implement in MacOS</p></td>
</tr>
</tbody>
</table>

## <span id="extended_attribute"></span>Extended attribute

### Extended attribute key data

The extended attribute key data (j\_xattr\_key\_t) is variable of size
and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>60 bits</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>File system object identifier (FSOID)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>7.4</p></td>
<td style="text-align: left;"><p>4 bits</p></td>
<td style="text-align: left;"><p>0x4</p></td>
<td style="text-align: left;"><p>File system data type<br />
See section: <a href="#file_system_data_types">File system data types</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Name string size (name_len)<br />
Includes the size of the end-of-string character</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>10</p></td>
<td style="text-align: left;"><p>…</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Name string (name)<br />
Contains an UTF-8 encoded string with an end-of-string character<br />
See section: <a href="#extended_attribute_names">Extended attribute names</a></p></td>
</tr>
</tbody>
</table>

### Extended attribute value data

The extended attribute value data (j\_xattr\_val\_t) is variable of size
and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Flags (flags)<br />
See section: <a href="#extended_attribute_flags">Extended attribute flags</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Extended attribute data size</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>…</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Extended attribute data</p></td>
</tr>
</tbody>
</table>

Extended attribute data size can contain 0 if extended attribute flags
XATTR\_DATA\_EMBEDDED is set.

### <span id="extended_attribute_names"></span>Extended attribute names

<table>
<colgroup>
<col style="width: 16%" />
<col style="width: 83%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Name</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>com.apple.assetsd.dbRebuildInProgress</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>com.apple.assetsd.dbRebuildUuid</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>com.apple.assetsd.thumbnailCameraPreviewImageAssetID</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>com.apple.assetsd.UUID</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>com.apple.decmpfs</p></td>
<td style="text-align: left;"><p>Compressed data extended attribute<br />
See section: <a href="#compressed_data_extended_attribute">Compressed data extended attribute</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>com.apple.FinderInfo</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>com.apple.fs.symlink</p></td>
<td style="text-align: left;"><p>Symbolic link</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>com.apple.genstore.info</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>com.apple.genstore.origdisplayname</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>com.apple.genstore.orig_perms_v1</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>com.apple.genstore.origposixname</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>com.apple.GeoServices.SHA1</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>com.apple.installd.installType</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>com.apple.installd.uniqueInstallID</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>com.apple.lastuseddate#PS</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>com.apple.metadata:_kMDItemUserTags</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>com.apple.metadata:com_apple_backup_excludeItem</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>com.apple.metadata:kMDItemDownloadedDate</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>com.apple.metadata:kMDItemWhereFroms</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>com.apple.metadata:kMDLabel_fwlfb7nbt2o7degof3q2o2btjy</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>com.apple.quarantine</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>com.apple.ResourceFork</p></td>
<td style="text-align: left;"><p>Resource fork</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>com.apple.rootless</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>com.apple.system.Security</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>com.apple.TextEncoding</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>LastUpgradeCheck</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>lock</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>org.chromium.crashpad.database.initialized</p></td>
<td style="text-align: left;"></td>
</tr>
</tbody>
</table>

### <span id="extended_attribute_flags"></span>Extended attribute flags

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0x0001</p></td>
<td style="text-align: left;"><p>XATTR_DATA_STREAM</p></td>
<td style="text-align: left;"><p>Extended attribute data is stored in a data stream<br />
The extended attribute data contains an 8-byte file system object identifier of the corresponding data stream<br />
See section: <a href="#extended_attribute_data_stream">Extended attribute data stream</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0002</p></td>
<td style="text-align: left;"><p>XATTR_DATA_EMBEDDED</p></td>
<td style="text-align: left;"><p>Extended attribute data is stored directly in the record</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x0004</p></td>
<td style="text-align: left;"><p>XATTR_FILE_SYSTEM_OWNED</p></td>
<td style="text-align: left;"><p>Extended attribute record is owned by the file system</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0008</p></td>
<td style="text-align: left;"><p>XATTR_RESERVED_8</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (Reserved)</span></strong></p></td>
</tr>
</tbody>
</table>

### <span id="extended_attribute_data_stream"></span>Extended attribute data stream

The extended attribute data stream (j\_xattr\_dstream\_t) is 48 bytes of
size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Data stream file system object identifier (xattr_obj_id)<br />
Contains the file system object identifier of the file extents that make up the data stream</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>48</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Data stream attribute<br />
See section: <a href="#data_stream_attribute">data stream attribute</a></p></td>
</tr>
</tbody>
</table>

### <span id="compressed_data_extended_attribute"></span>Compressed data extended attribute

The compressed extended attribute is named "com.apple.decmpfs" and
consists of:

-   compressed data header

-   optional compressed data

#### <span id="compressed_data_header"></span>Compressed data header

The compressed data header is 16 bytes of size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>"fpmc"</p></td>
<td style="text-align: left;"><p>Signature</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>Compression method<br />
See section: <a href="#compression_method">Compression method</a></p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
</tr>
</tbody>
</table>

The signature is likely stored in little-endian and represents "cmpf".

#### <span id="compression_method"></span>Compression method

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>1</p></td>
<td style="text-align: left;"><p>CMP_Type1</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (uncompressed extended attribute data)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>3</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>ZLIB (DEFLATE) compressed extended attribute data<br />
The compressed data is stored in the extended attribute after the compressed data header</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>64k chunked ZLIB (DEFLATE) compressed resource fork<br />
The compressed data is stored in the resource fork</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>5</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (sparse compressed extended attribute data)</span></strong><br />
Uncompressed data contains 0-byte values<br />
According to <code>[APPLE04]</code> specifies de-dup within the generation store.</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>6</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (unused)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>7</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>LZVN compressed extended attribute data<br />
The compressed data is stored in the extended attribute after the compressed data header</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>64k chunked LZVN compressed resource fork<br />
The compressed data is stored in the resource fork</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>9</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (uncompressed extended attribute data, different than CMP_Type1)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>10</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (64k chunked uncompressed data resource fork)</span></strong><br />
The compressed data is stored in the resource fork</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>11</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>LZFSE compressed extended attribute data<br />
The compressed data is stored in the extended attribute after the compressed data header</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>12</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>64k chunked LZFSE compressed resource fork<br />
The compressed data is stored in the resource fork</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x80000001</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (faulting file)</span></strong></p></td>
</tr>
</tbody>
</table>

If the ZLIB (DEFLATE) compressed data starts with 0xff the data is
stored uncompressed after the first compressed data byte. `[GANDER17]`
indicates that this should be `( byte_value & 0x0f ) == 0x0f`.

If the LZVN compressed data starts with 0x06 (end of stream oppcode) the
data is stored uncompressed after the first compressed data byte.

## <span id="sibling_link"></span>Sibling link

### Sibling link key data

The sibling link key data (j\_sibling\_key\_t) is 16 bytes of size and
consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>60 bits</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>File system object identifier (FSOID)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>7.4</p></td>
<td style="text-align: left;"><p>4 bits</p></td>
<td style="text-align: left;"><p>0x4</p></td>
<td style="text-align: left;"><p>File system data type<br />
See section: <a href="#file_system_data_types">File system data types</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Sibling map identifier (sibling_id)<br />
Contains the file system object identifier of the sibling map record</p></td>
</tr>
</tbody>
</table>

### Sibling link value data

The sibling link value data (j\_sibling\_val\_t) is variable of size and
consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Parent file system object identifier (parent_id)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Name string size (name_len)<br />
Includes the size of the end-of-string character</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>10</p></td>
<td style="text-align: left;"><p>…</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Name string (name)<br />
Contains an UTF-8 encoded string with an end-of-string character</p></td>
</tr>
</tbody>
</table>

## <span id="data_stream_identifier"></span>Data stream identifier

### Data stream identifier key data

The data stream key data (j\_dstream\_id\_key\_t) is 8 bytes of size and
consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>60 bits</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>File system object identifier (FSOID)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>7.4</p></td>
<td style="text-align: left;"><p>4 bits</p></td>
<td style="text-align: left;"><p>0x6</p></td>
<td style="text-align: left;"><p>File system data type<br />
See section: <a href="#file_system_data_types">File system data types</a></p></td>
</tr>
</tbody>
</table>

### Data stream identifier value data

The data stream value data (j\_dstream\_id\_val\_t) is 4 bytes of size
and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Reference count (refcnt)</p></td>
</tr>
</tbody>
</table>

## <span id="encryption_state"></span>Encryption state

**<span class="yellow-background">TODO: complete this section.</span>**

## <span id="file_extent"></span>File extent

### File extent key data

The file extent key data (j\_file\_extent\_key\_t) is 16 bytes of size
and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>60 bits</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>File system object identifier (FSOID)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>7.4</p></td>
<td style="text-align: left;"><p>4 bits</p></td>
<td style="text-align: left;"><p>0x8</p></td>
<td style="text-align: left;"><p>File system data type<br />
See section: <a href="#file_system_data_types">File system data types</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Logical address (logical_addr)<br />
Contains an offset relative to the start of the file entry data</p></td>
</tr>
</tbody>
</table>

### File extent value data

The file extent value data (j\_file\_extent\_val\_t) is 24 bytes of size
and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Extent data size and flags (len_and_flags)</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>7</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Extent data size</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>7</p></td>
<td style="text-align: left;"><p>1</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Flags<br />
See section: <a href="#file_extent_flags">File extent flags</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Physical block number (phys_block_num)<br />
Contains a block number relative to the start of the container</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Encryption identifier (crypto_id)<br />
Contains the <strong><span class="yellow-background">unknown</span></strong> and 0 if not set</p></td>
</tr>
</tbody>
</table>

### <span id="file_extent_flags"></span>File extent flags

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0x01</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (Is encrypted?)</span></strong></p></td>
</tr>
</tbody>
</table>

According to `[APPLE18]` there are currently no flags defined.
`[APPLE18]` also refers to `len_and_flags` as `len_and_kind`
interchangeably.

## <span id="directory_record"></span>Directory record

The directory record can have 2 different types of keys:

-   Key with name

-   Key with name and hash

It apprears that current APFS file system use a key with name and hash.
`[APPLE18]` does not indicate how to distinguish between the two, but
one method is to compare calculated and stored size of the key data.

In B-Tree branch nodes are sorted using the case-sensitive name, even
when the file system is case-insensitive.

### Directory record key data with name

The directory record key data with name (j\_drec\_key\_t) is variable of
size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object identifier and type (hdr)</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>60 bits</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>File system object identifier (FSOID)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>7.4</p></td>
<td style="text-align: left;"><p>4 bits</p></td>
<td style="text-align: left;"><p>0x9</p></td>
<td style="text-align: left;"><p>File system data type<br />
See section: <a href="#file_system_data_types">File system data types</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;" colspan="4"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Name string size (name_len)<br />
Includes the size of the end-of-string character</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>10</p></td>
<td style="text-align: left;"><p>…</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Name string (name)<br />
Contains an UTF-8 encoded string with an end-of-string character</p></td>
</tr>
</tbody>
</table>

### Directory record key data with name and hash

The directory record key data with name and hash
(j\_drec\_hashed\_key\_t) is variable of size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object identifier and type (hdr)</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>60 bits</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>File system object identifier (FSOID)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>7.4</p></td>
<td style="text-align: left;"><p>4 bits</p></td>
<td style="text-align: left;"><p>0x9</p></td>
<td style="text-align: left;"><p>File system data type<br />
See section: <a href="#file_system_data_types">File system data types</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;" colspan="4"><p><em>Name string size and hash (name_len_and_hash)</em></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>11 bits</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Name string size<br />
Includes the size of the end-of-string character</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>9.3</p></td>
<td style="text-align: left;"><p>21 bits</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Name hash<br />
See section: <a href="#directory_entry_name_hash">Directory entry name hash</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;" colspan="4"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>12</p></td>
<td style="text-align: left;"><p>…</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Name string (name)<br />
Contains an UTF-8 encoded string with an end-of-string character</p></td>
</tr>
</tbody>
</table>

### Directory record value data

The directory record value data (j\_drec\_val\_t) is variable of size
and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>File system object identifier of the directory entry (file_id)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Date and time the directory entry was added (date_added)<br />
Signed integer that contains the number of nanoseconds since January 1, 1970 00:00:00 UTC or 0 if not set</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Directory entry flags<br />
See section: <a href="#directory_entry_flags">Directory entry flags</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>18</p></td>
<td style="text-align: left;"><p>…</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Extended fields (xfields)<br />
See section: <a href="#extended_fields">Extended fields</a></p></td>
</tr>
</tbody>
</table>

#### <span id="directory_entry_flags"></span>Directory entry flags

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0x0000</p></td>
<td style="text-align: left;"><p>DT_UNKNOWN</p></td>
<td style="text-align: left;"><p>Unknown</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0001</p></td>
<td style="text-align: left;"><p>DT_FIFO</p></td>
<td style="text-align: left;"><p>Named pipe</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x0002</p></td>
<td style="text-align: left;"><p>DT_CHR</p></td>
<td style="text-align: left;"><p>Character-special file (Character device)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x0004</p></td>
<td style="text-align: left;"><p>DT_DIR</p></td>
<td style="text-align: left;"><p>Directory</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x0006</p></td>
<td style="text-align: left;"><p>DT_BLK</p></td>
<td style="text-align: left;"><p>Block-special file (Block device)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x0008</p></td>
<td style="text-align: left;"><p>DT_REG</p></td>
<td style="text-align: left;"><p>Regular file</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x000a</p></td>
<td style="text-align: left;"><p>DT_LNK</p></td>
<td style="text-align: left;"><p>Symbolic link</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x000c</p></td>
<td style="text-align: left;"><p>DT_SOCK</p></td>
<td style="text-align: left;"><p>Socket</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x000e</p></td>
<td style="text-align: left;"><p>DT_WHT</p></td>
<td style="text-align: left;"><p>Whiteout<br />
A whiteout is a directory entry that covers up all entries of a particular name from lower branches</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x000f</p></td>
<td style="text-align: left;"><p>DREC_TYPE_MASK</p></td>
<td style="text-align: left;"><p>Directory type bitmask</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x0010</p></td>
<td style="text-align: left;"><p>RESERVED_10</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (reserved)</span></strong></p></td>
</tr>
</tbody>
</table>

#### <span id="directory_entry_name_hash"></span>Directory entry name hash

The name hash of a directory entry is calculated as following:

-   If the file system is case-insensitive represent the name in
    lower-case

-   Represent the name as an Unicode string in Normalization Form
    Canonical Decomposition (NFD)

-   Format the Unicode string as a little-endian UTF-32 stream without a
    byte-order-mark or end-of-string character

-   Calculate a CRC-32c checksum of the UTF-32 stream with an initial
    checkum of 0xffffffff (-1)

-   The lower 22-bits of checksum form the hash

The CRC-32 calculation uses the Castagnoli polynomial of 0x1edc6f41,
also known as CRC-32C. The CRC-32 calculation does not use the XOR with
0xffffffff before and after the calculation, which is also referred to
as weak CRC-32 calculation.

## <span id="directory_stats"></span>Directory stats

### Directory stats key data

The directory stats key data (j\_dir\_stats\_key\_t) is 8 bytes of size
and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>60 bits</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>File system object identifier (FSOID)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>7.4</p></td>
<td style="text-align: left;"><p>4 bits</p></td>
<td style="text-align: left;"><p>0xa</p></td>
<td style="text-align: left;"><p>File system data type<br />
See section: <a href="#file_system_data_types">File system data types</a></p></td>
</tr>
</tbody>
</table>

### Directory stats value data

The directory stats value data (j\_dir\_stats\_val\_t) is 32 bytes of
size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of children (num_children)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Total size (total_size)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Parent directory file system object identifier (chained_key)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Generation count (gen_count)</p></td>
</tr>
</tbody>
</table>

## <span id="snapshot_name"></span>Snapshot name

The snapshot name (j\_snap\_name\_val\_t) is 8 bytes of size and
consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>60 bits</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Snapshot metdata object identifier</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>7.4</p></td>
<td style="text-align: left;"><p>4 bits</p></td>
<td style="text-align: left;"><p>0x1</p></td>
<td style="text-align: left;"><p>File system data type<br />
See section: <a href="#file_system_data_types">File system data types</a></p></td>
</tr>
</tbody>
</table>

## <span id="sibling_map"></span>Sibling map

### Sibling map key data

The sibling map key data (j\_sibling\_map\_key\_t) is 8 bytes of size
and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>60 bits</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>File system object identifier (FSOID)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>7.4</p></td>
<td style="text-align: left;"><p>4 bits</p></td>
<td style="text-align: left;"><p>0x4</p></td>
<td style="text-align: left;"><p>File system data type<br />
See section: <a href="#file_system_data_types">File system data types</a></p></td>
</tr>
</tbody>
</table>

### Sibling map value data

The sibling map value data (j\_sibling\_map\_val\_t) is 8 bytes of size
and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>File system object identifier (file_id)</p></td>
</tr>
</tbody>
</table>

## <span id="extended_fields"></span>Extended fields

Directory entries and inodes use extended fields to store additional
attributes, such as the filename.

The extended fields (xf\_blob\_t) consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of extended fields (xf_num_exts)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Extended field value data size (xf_used_data)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Extended field data (xf_data)</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>…</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Array of extended field descriptors<br />
See section: <a href="#extended_field_descriptor">Extended field descriptor</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>…</p></td>
<td style="text-align: left;"><p>…</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Extended field value data</p></td>
</tr>
</tbody>
</table>

The extended field values are stored 8-byte aligned in the extended
field value data.

### <span id="extended_field_descriptor"></span>Extended field descriptor

An extended field descriptor (x\_field\_t) is 4 bytes of size and
consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>1</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Extended field type (x_type)<br />
See section: <a href="#extended_field_types">Extended field types</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>1</p></td>
<td style="text-align: left;"><p>1</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Extended field flags (x_flags)<br />
See section: <a href="#extended_field_flags">Extended field flags</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Extended field data size (x_size)</p></td>
</tr>
</tbody>
</table>

### <span id="extended_field_types"></span>Extended field types

#### Directory record extended field types

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>1</p></td>
<td style="text-align: left;"><p>DREC_EXT_TYPE_SIBLING_ID</p></td>
<td style="text-align: left;"><p>Hard link sibling identifier<br />
The extended field data contains a 64-bit integer value</p></td>
</tr>
</tbody>
</table>

#### <span id="inode_extended_field_types"></span>Inode extended field types

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>1</p></td>
<td style="text-align: left;"><p>INO_EXT_TYPE_SNAP_XID</p></td>
<td style="text-align: left;"><p>Transaction identifier of a snapshot<br />
The extended field data contains a 64-bit integer value</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"><p>INO_EXT_TYPE_DELTA_TREE_OID</p></td>
<td style="text-align: left;"><p>Object identifier of the snapshot extent delta list<br />
The extended field data contains a 64-bit integer value</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>3</p></td>
<td style="text-align: left;"><p>INO_EXT_TYPE_DOCUMENT_ID</p></td>
<td style="text-align: left;"><p>Document identifier<br />
The extended field data contains a 32-bit integer value</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>INO_EXT_TYPE_NAME</p></td>
<td style="text-align: left;"><p>Filename<br />
The extended field data contains an UTF-8 string with end-of-string character</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>5</p></td>
<td style="text-align: left;"><p>INO_EXT_TYPE_PREV_FSIZE</p></td>
<td style="text-align: left;"><p>Previous file size<br />
The extended field data contains a 64-bit integer value</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>6</p></td>
<td style="text-align: left;"><p>INO_EXT_TYPE_RESERVED_6</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (Reserved)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>7</p></td>
<td style="text-align: left;"><p>INO_EXT_TYPE_FINDER_INFO</p></td>
<td style="text-align: left;"><p>Finder information<br />
The extended field data contains a 32-bit integer value</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>INO_EXT_TYPE_DSTREAM</p></td>
<td style="text-align: left;"><p>Data stream<br />
The extended field data contains a <a href="#data_stream_attribute">data stream attribute</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>9</p></td>
<td style="text-align: left;"><p>INO_EXT_TYPE_RESERVED_9</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (Reserved)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>10</p></td>
<td style="text-align: left;"><p>INO_EXT_TYPE_DIR_STATS_KEY</p></td>
<td style="text-align: left;"><p>Direcotry statistics<br />
<strong><span class="yellow-background">Unknown if this contains the object identifier of the directory statisticts or a j_dir_stats_val_t structure</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>11</p></td>
<td style="text-align: left;"><p>INO_EXT_TYPE_FS_UUID</p></td>
<td style="text-align: left;"><p>Mounted file system identifier<br />
The extended field data contains a 128-bit UUID value</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>12</p></td>
<td style="text-align: left;"><p>INO_EXT_TYPE_RESERVED_12</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (Reserved)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>13</p></td>
<td style="text-align: left;"><p>INO_EXT_TYPE_SPARSE_BYTES</p></td>
<td style="text-align: left;"><p>Number of sparse bytes in the data stream<br />
The extended field data contains a 64-bit integer value</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>14</p></td>
<td style="text-align: left;"><p>INO_EXT_TYPE_RDEV</p></td>
<td style="text-align: left;"><p>Block- or character-device identifier<br />
The extended field data contains a 32-bit integer value</p></td>
</tr>
</tbody>
</table>

### <span id="extended_field_flags"></span>Extended field flags

<table>
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 71%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Identifier</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0x01</p></td>
<td style="text-align: left;"><p>XF_DATA_DEPENDENT</p></td>
<td style="text-align: left;"><p>Contents of the extended field is dependent on the data stream (file contents)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x02</p></td>
<td style="text-align: left;"><p>XF_DO_NOT_COPY</p></td>
<td style="text-align: left;"><p>Do not duplicate the extended field when copied</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x04</p></td>
<td style="text-align: left;"><p>XF_RESERVED_4</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (Reserved)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x08</p></td>
<td style="text-align: left;"><p>XF_CHILDREN_INHERIT</p></td>
<td style="text-align: left;"><p>Newly created sub directory entries (children) inherit the extended field</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x10</p></td>
<td style="text-align: left;"><p>XF_USER_FIELD</p></td>
<td style="text-align: left;"><p>Extended field was added by an user-space program</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x20</p></td>
<td style="text-align: left;"><p>XF_SYSTEM_FIELD</p></td>
<td style="text-align: left;"><p>Extended field was added by the system (kernel)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>0x40</p></td>
<td style="text-align: left;"><p>XF_RESERVED_40</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (Reserved)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0x80</p></td>
<td style="text-align: left;"><p>XF_RESERVED_80</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (Reserved)</span></strong></p></td>
</tr>
</tbody>
</table>

## <span id="data_stream_attribute"></span>Data stream attribute

The data stream attribute (j\_dstream\_t) is 40 bytes of size and
consist of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Used size (size)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Allocated size (alloced_size)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>(Default) encryption identifier (default_crypto_id)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Total number of bytes written to data stream (total_bytes_written)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>32</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Total number of bytes read from data stream (total_bytes_written)</p></td>
</tr>
</tbody>
</table>

## <span id="file_content"></span>File content

APFS supports multiple ways to store file content:

-   Data fork

-   Compressed data extended attribute

-   Compressed data extended attribute with resource fork

-   Resource fork

-   Extended attribute (named fork)

## Data fork

The file content size is stored in an INO\_EXT\_TYPE\_DSTREAM [inode
extended field type](#inode_extended_field_types).

The file content data can be located through the [file
extents](#file_extent) for the data stream file system object identifier
in the [file system tree](#file_system).

If the volume is encrypted the file content is encrypted with the
encryption identifier in defined by the [File extent](#file_extent).

If the [inode flag](#inode_flags) INODE\_IS\_SPARSE is set the file
contains one or more spare file extents. A sparse file extent has a
physical block number of 0.

## Compressed data extended attribute

[Compression method](#compression_method) should be 3, 5 or 7.

The file content size is stored in the compressed data header of a
"com.apple.decmpfs" [extended attribute](#extended_attribute).

For [compression method](#compression_method) 3 or 7 the file content
data is stored in a "com.apple.decmpfs" [extended
attribute](#extended_attribute) after the [compressed data
header](#compressed_data_header).

For [compression method](#compression_method) 5 the file content data
contains 0-byte values. There are 12 bytes stored after the [compressed
data header](#compressed_data_header) that contain:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong><br />
Seen: 1</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong><br />
Seen: 0</p></td>
</tr>
</tbody>
</table>

## Compressed data extended attribute with resource fork

[Compression method](#compression_method) should be 4 or 8.

The file content size is stored in the compressed data header of a
"com.apple.decmpfs" [extended attribute](#extended_attribute).

The file content data is stored in a "com.apple.ResourceFork" [extended
attribute](#extended_attribute).

The compressed data starts with metadata that contains the offsets of
the compressed data blocks.

### ZLIB (DEFLATE) compressed data

-   ZLIB (DEFLATE) compressed header

-   **<span class="yellow-background">Unknown (empty values)</span>**

-   ZLIB (DEFLATE) compressed data block offsets and sizes

-   ZLIB (DEFLATE) compressed data blocks

-   ZLIB (DEFLATE) compressed footer

#### ZLIB (DEFLATE) compressed header

The ZLIB (DEFLATE) compressed header is 16 bytes of size and consists
of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Compressed data block descriptors offset<br />
The offset is relative from the start of the ZLIB (DEFLATE) compressed data</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Compressed footer offset<br />
The offset is relative from the start of the ZLIB (DEFLATE) compressed data</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Compressed data block descriptors and data size</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>12</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Compressed footer size</p></td>
</tr>
</tbody>
</table>

The values in the ZLIB (DEFLATE) compressed header are stored in
big-endian.

#### ZLIB (DEFLATE) compressed data block descriptors

The ZLIB (DEFLATE) compressed data block descriptors are variable in
size and consist of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Compressed data size</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of compressed data block offset and size tuples</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8 x …</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Array of compressed data block descriptors</p></td>
</tr>
</tbody>
</table>

#### ZLIB (DEFLATE) compressed data block descriptor

The ZLIB (DEFLATE) compressed data block descriptor is 8 bytes of size
and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Compressed block offset<br />
The offset is relative from the start of the ZLIB (DEFLATE) compressed data + 20</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Compressed block size</p></td>
</tr>
</tbody>
</table>

#### ZLIB (DEFLATE) compressed footer

The ZLIB (DEFLATE) compressed footer is 50 bytes size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (empty values)</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>26</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>28</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>30</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>32</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>"cmpf"</p></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (signature)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>36</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>40</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>44</p></td>
<td style="text-align: left;"><p>6</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (empty values)</span></strong></p></td>
</tr>
</tbody>
</table>

The values in the ZLIB (DEFLATE) compressed footer are stored in
big-endian.

### LZVN compressed data

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>4 x …</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Array of compressed data block offsets<br />
The offset is relative from the start of the LZVN compressed data</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>…</p></td>
<td style="text-align: left;"><p>…</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>LZVN compressed data blocks</p></td>
</tr>
</tbody>
</table>

The compressed data block contains a maximum of 65536 bytes of data. The
compressed data block therefore should not exceed 65537 bytes of size.

## Resource fork

**<span class="yellow-background">TODO: complete this section.</span>**

## Extended attribute (named fork)

**<span class="yellow-background">TODO: complete this section.</span>**

## <span id="efi_jumpstart"></span>EFI jumpstart

The EFI jumpstart (nx\_efi\_jumpstart\_t) is variable of size and
consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object header</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object checksum<br />
See section: <a href="#object_checksum">Object checkum</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object identifier</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object transaction identifier (xid)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x00000014</p></td>
<td style="text-align: left;"><p>Object type<br />
See section: <a href="#object_types">Object types</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>28</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x00000000</p></td>
<td style="text-align: left;"><p>Object subtype</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object values</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>32</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>"RDSJ"</p></td>
<td style="text-align: left;"><p>Signature (nej_magic)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>36</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>1</p></td>
<td style="text-align: left;"><p>Format version (nej_version)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>40</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (nej_efi_file_len?)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>44</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of extents (nej_num_extents)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>48</p></td>
<td style="text-align: left;"><p>16 x 8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown (nej_reserved?)</span></strong></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>176</p></td>
<td style="text-align: left;"><p>number of extents x 16</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Extents (nej_rec_extents)<br />
Contains the extents where the EFI driver is stored<br />
See section: <a href="#efi_jumpstart_extent">EFI jumpstart extent</a></p></td>
</tr>
</tbody>
</table>

## <span id="efi_jumpstart_extent"></span>EFI jumpstart extent

The EFI jumpstart extent (prange\_t) is 16 bytes of size and consists
of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Block number</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Number of blocks</p></td>
</tr>
</tbody>
</table>

## <span id="extent_reference_tree"></span>Extent-reference tree

**<span class="yellow-background">TODO: complete this section.</span>**

## Snapshots

**<span class="yellow-background">TODO: complete this section.</span>**

## <span id="snapshot_metadata_tree"></span>Snapshot metadata tree

The snapshot metadata tree consists of:

-   snapshot metadata tree (object)

-   snapshot metadata B-tree

## Snapshot metadata tree object

The snapshot metadata tree object is 32 bytes of size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object header</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object checksum<br />
See section: <a href="#object_checksum">Object checkum</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object identifier</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object transaction identifier (xid)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x40000002<br />
0x40000003</p></td>
<td style="text-align: left;"><p>Object type<br />
See section: <a href="#object_types">Object types</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>28</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x00000010</p></td>
<td style="text-align: left;"><p>Object subtype</p></td>
</tr>
</tbody>
</table>

## Snapshot metadata B-tree

The object map values are stored in [B-tree](#btree).

### Snapshot metadata B-tree key

The snapshot metadata B-tree key (j\_snap\_metadata\_key\_t or
j\_snap\_name\_key\_t) is variable of size and consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Key object identifier (hdr)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;" colspan="4"><p><em>If key object identifier data type is APFS_TYPE_SNAP_NAME</em></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>…</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Snapshot name string<br />
Contains an UTF-8 encoded string with an end-of-string character</p></td>
</tr>
</tbody>
</table>

### Snapshot metadata B-tree branch node value

An snapshot metadata B-tree node contains branch node values if
BTNODE\_LEAF is not set. The corresponding inapshot metadata B-tree key
represents the first key in the branch.

An napshot metadata B-tree branch node value is 8 bytes of size and
consists of:

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Sub node block number<br />
Contains a block number relative to the start of the container</p></td>
</tr>
</tbody>
</table>

### Snapshot metadata B-tree leaf node value

The contents of a snapshot metadata B-tree leaf node depends on the
[File system data type](#file_system_data_types) of the key object
identifier.

<table>
<colgroup>
<col style="width: 16%" />
<col style="width: 83%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>APFS_TYPE_SNAP_METADATA</p></td>
<td style="text-align: left;"><p>Snapshot object identifier<br />
The value contains <a href="#snapshot_metadata">Snapshot metadata</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>APFS_TYPE_SNAP_NAME</p></td>
<td style="text-align: left;"><p>Snapshot name<br />
The value contains <a href="#snapshot_name">Snapshot name</a></p></td>
</tr>
</tbody>
</table>

## Fusion drives

A Fusion drive consists of a main SSD and a tier2 magnetic disk that
together form one logical APFS container.

## <span id="fusion_middle_tree"></span>Fusion middle tree

**<span class="yellow-background">TODO: complete this section.</span>**

<table>
<colgroup>
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 12%" />
<col style="width: 63%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Offset</th>
<th style="text-align: left;">Size</th>
<th style="text-align: left;">Value</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object header</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object checksum<br />
See section: <a href="#object_checksum">Object checkum</a></p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object identifier</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>16</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Object transaction identifier (xid)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>24</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x40000002</p></td>
<td style="text-align: left;"><p>Object type<br />
See section: <a href="#object_types">Object types</a></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>28</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>0x00000015</p></td>
<td style="text-align: left;"><p>Object subtype</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;" colspan="4"><p><em>Object values</em></p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>…</p></td>
<td style="text-align: left;"><p>…</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p><strong><span class="yellow-background">Unknown</span></strong></p></td>
</tr>
</tbody>
</table>

## Notes

TODO describe evict\_mapping\_val\_t

    fsck_apfs -n apfs_empty.raw
    ** Checking volume.
    ** Checking the container superblock.
    ** Checking the space manager.
    ** Checking the object map.
    ** Checking the APFS volume superblock.
    ** Checking the object map.
    mount_apfs: mount: Operation not supported by device
    error: mount_apfs exit status 73
    ** Checking the fsroot tree.
    ** Checking the snapshot metadata tree.
    ** Checking the extent ref tree.
    ** Checking the snapshots.
    warning: unmount: /private/var/folders/7k/s4ykb4095ld38gz50kvkhf_m0000gn/T/fsck_apfs.586: Invalid argument
    ** Verifying allocated space.
    ** The volume apfs_empty.raw appears to be OK.

    diskutil apfs list
    APFS Container (1 found)
    |
    +-- Container disk4 B069A33C-65E6-4DAF-BD60-081493F91116
        ====================================================
        APFS Container Reference:     disk4
        Size (Capacity Ceiling):      1007616 B (1.0 MB)
        Minimum Size:                 1007616 B (1.0 MB)
        Capacity In Use By Volumes:   376832 B (376.8 KB) (37.4% used)
        Capacity Not Allocated:       630784 B (630.8 KB) (62.6% free)
        |
        +-< Physical Store disk3s1 D05BEA77-81B3-4A8D-A82A-BC5C7D2FB27C
        |   -----------------------------------------------------------
        |   APFS Physical Store Disk:   disk3s1
        |   Size:                       1007616 B (1.0 MB)
        |
        +-> Volume disk4s1 9428213B-2E82-49E9-871E-7220B157FC8D
            ---------------------------------------------------
            APFS Volume Disk (Role):   disk4s1 (No specific role)
            Name:                      untitled (Case-insensitive)
            Mount Point:               /Volumes/untitled
            Capacity Consumed:         24576 B (24.6 KB)
            FileVault:                 No

    find /Volumes/SingleVolume
    /Volumes/SingleVolume
    /Volumes/SingleVolume/.fseventsd
    /Volumes/SingleVolume/.fseventsd/fseventsd-uuid
    /Volumes/SingleVolume/.fseventsd/0000000000793354
    /Volumes/SingleVolume/.fseventsd/0000000000793353

decmpfs is also referred to as AFSC (Apple File System Compression) or
HFS/HFS+ compression

## References

`[APPLE04]`

<table>
<colgroup>
<col style="width: 16%" />
<col style="width: 83%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Title:</th>
<th style="text-align: left;">copyfile.c</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>Author(s):</p></td>
<td style="text-align: left;"><p>Apple</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Date:</p></td>
<td style="text-align: left;"><p>2004</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>URL:</p></td>
<td style="text-align: left;"><p><a href="https://opensource.apple.com/source/copyfile/copyfile-127/copyfile.c">https://opensource.apple.com/source/copyfile/copyfile-127/copyfile.c</a></p></td>
</tr>
</tbody>
</table>

`[APPLE18]`

<table>
<colgroup>
<col style="width: 16%" />
<col style="width: 83%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Title:</th>
<th style="text-align: left;">Apple File System Reference</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>Author(s):</p></td>
<td style="text-align: left;"><p>Apple</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Date:</p></td>
<td style="text-align: left;"><p>September 2018</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>URL:</p></td>
<td style="text-align: left;"><p><a href="https://developer.apple.com/support/apple-file-system/Apple-File-System-Reference.pdf">https://developer.apple.com/support/apple-file-system/Apple-File-System-Reference.pdf</a></p></td>
</tr>
</tbody>
</table>

`[GANDER17]`

<table>
<colgroup>
<col style="width: 16%" />
<col style="width: 83%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Title:</th>
<th style="text-align: left;">apfs_fuse</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>Author(s):</p></td>
<td style="text-align: left;"><p>S. Gander</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Date:</p></td>
<td style="text-align: left;"><p>October 2017</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>URL:</p></td>
<td style="text-align: left;"><p><a href="https://github.com/sgan81/apfs-fuse">https://github.com/sgan81/apfs-fuse</a></p></td>
</tr>
</tbody>
</table>

`[IEEE 1619-2007]`

<table>
<colgroup>
<col style="width: 16%" />
<col style="width: 83%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Title:</th>
<th style="text-align: left;">The XTS-AES Tweakable Block Cipher (IEEE 1619-2007)</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>Author(s):</p></td>
<td style="text-align: left;"><p>IEEE</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Date:</p></td>
<td style="text-align: left;"><p>April 18, 2008</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>URL:</p></td>
<td style="text-align: left;"><p><a href="http://axelkenzo.ru/downloads/1619-2007-NIST-Submission.pdf">http://axelkenzo.ru/downloads/1619-2007-NIST-Submission.pdf</a><br />
<a href="https://bitbucket.org/garethl/xtssharp/src/0e6a81a823e9/docs/1619-2007-NIST-Submission.pdf">https://bitbucket.org/garethl/xtssharp/src/0e6a81a823e9/docs/1619-2007-NIST-Submission.pdf</a></p></td>
</tr>
</tbody>
</table>

`[HANSEN17]`

<table>
<colgroup>
<col style="width: 16%" />
<col style="width: 83%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Title:</th>
<th style="text-align: left;">Decoding the APFS file system</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>Author(s):</p></td>
<td style="text-align: left;"><p>Hansen, Kurt &amp; Toolan, Fergus</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Date:</p></td>
<td style="text-align: left;"><p>September 2017</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>URL:</p></td>
<td style="text-align: left;"><p><a href="https://www.researchgate.net/publication/319573636_Decoding_the_APFS_file_system">https://www.researchgate.net/publication/319573636_Decoding_the_APFS_file_system</a></p></td>
</tr>
</tbody>
</table>

`[KODIS92]`

<table>
<colgroup>
<col style="width: 16%" />
<col style="width: 83%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Title:</th>
<th style="text-align: left;">Fletcher’s checksum - Error correction at a fraction of the cost</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>Author(s):</p></td>
<td style="text-align: left;"><p>John Kodis</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Date:</p></td>
<td style="text-align: left;"><p>May 1992</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>URL:</p></td>
<td style="text-align: left;"><p><a href="http://collaboration.cmc.ec.gc.ca/science/rpn/biblio/ddj/Website/articles/DDJ/1992/9205/9205b/9205b.htm">http://collaboration.cmc.ec.gc.ca/science/rpn/biblio/ddj/Website/articles/DDJ/1992/9205/9205b/9205b.htm</a></p></td>
</tr>
</tbody>
</table>

`[PLUM17]`

<table>
<colgroup>
<col style="width: 16%" />
<col style="width: 83%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Title:</th>
<th style="text-align: left;">APFS filesystem format</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>Author(s):</p></td>
<td style="text-align: left;"><p>Jonas Plum</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Date:</p></td>
<td style="text-align: left;"><p>April 2017</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>URL:</p></td>
<td style="text-align: left;"><p><a href="https://blog.cugu.eu/post/apfs/">https://blog.cugu.eu/post/apfs/</a></p></td>
</tr>
</tbody>
</table>

`[RFC2898]`

<table>
<colgroup>
<col style="width: 16%" />
<col style="width: 83%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Title:</th>
<th style="text-align: left;">PKCS #5: Password-Based Cryptography Specification</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>Version:</p></td>
<td style="text-align: left;"><p>2.0</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Author(s):</p></td>
<td style="text-align: left;"><p>B. Kaliski</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>Date:</p></td>
<td style="text-align: left;"><p>September 2000</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>URL:</p></td>
<td style="text-align: left;"><p><a href="https://www.ietf.org/rfc/rfc2898.txt">https://www.ietf.org/rfc/rfc2898.txt</a></p></td>
</tr>
</tbody>
</table>

`[RFC3394]`

<table>
<colgroup>
<col style="width: 16%" />
<col style="width: 83%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Title:</th>
<th style="text-align: left;">Advanced Encryption Standard (AES) Key Wrap Algorithm</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>Author(s):</p></td>
<td style="text-align: left;"><p>J. Schaad, R. Housley</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Date:</p></td>
<td style="text-align: left;"><p>September 2002</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>URL:</p></td>
<td style="text-align: left;"><p><a href="https://www.ietf.org/rfc/rfc3394.txt">https://www.ietf.org/rfc/rfc3394.txt</a></p></td>
</tr>
</tbody>
</table>

// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto/manufacturer.proto
// Protobuf Java Version: 4.26.1

package com.nyrta1.SkyTicket_AdminService.models;

public final class ManufacturerEntity {
  private ManufacturerEntity() {}
  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
      com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
      /* major= */ 4,
      /* minor= */ 26,
      /* patch= */ 1,
      /* suffix= */ "",
      ManufacturerEntity.class.getName());
  }
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public interface ManufacturerOrBuilder extends
      // @@protoc_insertion_point(interface_extends:Manufacturer)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <code>int64 id = 1;</code>
     * @return The id.
     */
    long getId();

    /**
     * <code>string name = 2;</code>
     * @return The name.
     */
    java.lang.String getName();
    /**
     * <code>string name = 2;</code>
     * @return The bytes for name.
     */
    com.google.protobuf.ByteString
        getNameBytes();

    /**
     * <code>.google.protobuf.Timestamp created_at = 3;</code>
     * @return Whether the createdAt field is set.
     */
    boolean hasCreatedAt();
    /**
     * <code>.google.protobuf.Timestamp created_at = 3;</code>
     * @return The createdAt.
     */
    com.google.protobuf.Timestamp getCreatedAt();
    /**
     * <code>.google.protobuf.Timestamp created_at = 3;</code>
     */
    com.google.protobuf.TimestampOrBuilder getCreatedAtOrBuilder();
  }
  /**
   * Protobuf type {@code Manufacturer}
   */
  public static final class Manufacturer extends
      com.google.protobuf.GeneratedMessage implements
      // @@protoc_insertion_point(message_implements:Manufacturer)
      ManufacturerOrBuilder {
  private static final long serialVersionUID = 0L;
    static {
      com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
        com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
        /* major= */ 4,
        /* minor= */ 26,
        /* patch= */ 1,
        /* suffix= */ "",
        Manufacturer.class.getName());
    }
    // Use Manufacturer.newBuilder() to construct.
    private Manufacturer(com.google.protobuf.GeneratedMessage.Builder<?> builder) {
      super(builder);
    }
    private Manufacturer() {
      name_ = "";
    }

    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.internal_static_Manufacturer_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessage.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.internal_static_Manufacturer_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer.class, com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer.Builder.class);
    }

    private int bitField0_;
    public static final int ID_FIELD_NUMBER = 1;
    private long id_ = 0L;
    /**
     * <code>int64 id = 1;</code>
     * @return The id.
     */
    @java.lang.Override
    public long getId() {
      return id_;
    }

    public static final int NAME_FIELD_NUMBER = 2;
    @SuppressWarnings("serial")
    private volatile java.lang.Object name_ = "";
    /**
     * <code>string name = 2;</code>
     * @return The name.
     */
    @java.lang.Override
    public java.lang.String getName() {
      java.lang.Object ref = name_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        name_ = s;
        return s;
      }
    }
    /**
     * <code>string name = 2;</code>
     * @return The bytes for name.
     */
    @java.lang.Override
    public com.google.protobuf.ByteString
        getNameBytes() {
      java.lang.Object ref = name_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        name_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static final int CREATED_AT_FIELD_NUMBER = 3;
    private com.google.protobuf.Timestamp createdAt_;
    /**
     * <code>.google.protobuf.Timestamp created_at = 3;</code>
     * @return Whether the createdAt field is set.
     */
    @java.lang.Override
    public boolean hasCreatedAt() {
      return ((bitField0_ & 0x00000001) != 0);
    }
    /**
     * <code>.google.protobuf.Timestamp created_at = 3;</code>
     * @return The createdAt.
     */
    @java.lang.Override
    public com.google.protobuf.Timestamp getCreatedAt() {
      return createdAt_ == null ? com.google.protobuf.Timestamp.getDefaultInstance() : createdAt_;
    }
    /**
     * <code>.google.protobuf.Timestamp created_at = 3;</code>
     */
    @java.lang.Override
    public com.google.protobuf.TimestampOrBuilder getCreatedAtOrBuilder() {
      return createdAt_ == null ? com.google.protobuf.Timestamp.getDefaultInstance() : createdAt_;
    }

    private byte memoizedIsInitialized = -1;
    @java.lang.Override
    public final boolean isInitialized() {
      byte isInitialized = memoizedIsInitialized;
      if (isInitialized == 1) return true;
      if (isInitialized == 0) return false;

      memoizedIsInitialized = 1;
      return true;
    }

    @java.lang.Override
    public void writeTo(com.google.protobuf.CodedOutputStream output)
                        throws java.io.IOException {
      if (id_ != 0L) {
        output.writeInt64(1, id_);
      }
      if (!com.google.protobuf.GeneratedMessage.isStringEmpty(name_)) {
        com.google.protobuf.GeneratedMessage.writeString(output, 2, name_);
      }
      if (((bitField0_ & 0x00000001) != 0)) {
        output.writeMessage(3, getCreatedAt());
      }
      getUnknownFields().writeTo(output);
    }

    @java.lang.Override
    public int getSerializedSize() {
      int size = memoizedSize;
      if (size != -1) return size;

      size = 0;
      if (id_ != 0L) {
        size += com.google.protobuf.CodedOutputStream
          .computeInt64Size(1, id_);
      }
      if (!com.google.protobuf.GeneratedMessage.isStringEmpty(name_)) {
        size += com.google.protobuf.GeneratedMessage.computeStringSize(2, name_);
      }
      if (((bitField0_ & 0x00000001) != 0)) {
        size += com.google.protobuf.CodedOutputStream
          .computeMessageSize(3, getCreatedAt());
      }
      size += getUnknownFields().getSerializedSize();
      memoizedSize = size;
      return size;
    }

    @java.lang.Override
    public boolean equals(final java.lang.Object obj) {
      if (obj == this) {
       return true;
      }
      if (!(obj instanceof com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer)) {
        return super.equals(obj);
      }
      com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer other = (com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer) obj;

      if (getId()
          != other.getId()) return false;
      if (!getName()
          .equals(other.getName())) return false;
      if (hasCreatedAt() != other.hasCreatedAt()) return false;
      if (hasCreatedAt()) {
        if (!getCreatedAt()
            .equals(other.getCreatedAt())) return false;
      }
      if (!getUnknownFields().equals(other.getUnknownFields())) return false;
      return true;
    }

    @java.lang.Override
    public int hashCode() {
      if (memoizedHashCode != 0) {
        return memoizedHashCode;
      }
      int hash = 41;
      hash = (19 * hash) + getDescriptor().hashCode();
      hash = (37 * hash) + ID_FIELD_NUMBER;
      hash = (53 * hash) + com.google.protobuf.Internal.hashLong(
          getId());
      hash = (37 * hash) + NAME_FIELD_NUMBER;
      hash = (53 * hash) + getName().hashCode();
      if (hasCreatedAt()) {
        hash = (37 * hash) + CREATED_AT_FIELD_NUMBER;
        hash = (53 * hash) + getCreatedAt().hashCode();
      }
      hash = (29 * hash) + getUnknownFields().hashCode();
      memoizedHashCode = hash;
      return hash;
    }

    public static com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseWithIOException(PARSER, input);
    }
    public static com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseWithIOException(PARSER, input, extensionRegistry);
    }

    public static com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseDelimitedWithIOException(PARSER, input);
    }

    public static com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseWithIOException(PARSER, input);
    }
    public static com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer parseFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseWithIOException(PARSER, input, extensionRegistry);
    }

    @java.lang.Override
    public Builder newBuilderForType() { return newBuilder(); }
    public static Builder newBuilder() {
      return DEFAULT_INSTANCE.toBuilder();
    }
    public static Builder newBuilder(com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer prototype) {
      return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
    }
    @java.lang.Override
    public Builder toBuilder() {
      return this == DEFAULT_INSTANCE
          ? new Builder() : new Builder().mergeFrom(this);
    }

    @java.lang.Override
    protected Builder newBuilderForType(
        com.google.protobuf.GeneratedMessage.BuilderParent parent) {
      Builder builder = new Builder(parent);
      return builder;
    }
    /**
     * Protobuf type {@code Manufacturer}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessage.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:Manufacturer)
        com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.ManufacturerOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.internal_static_Manufacturer_descriptor;
      }

      @java.lang.Override
      protected com.google.protobuf.GeneratedMessage.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.internal_static_Manufacturer_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer.class, com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer.Builder.class);
      }

      // Construct using com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer.newBuilder()
      private Builder() {
        maybeForceBuilderInitialization();
      }

      private Builder(
          com.google.protobuf.GeneratedMessage.BuilderParent parent) {
        super(parent);
        maybeForceBuilderInitialization();
      }
      private void maybeForceBuilderInitialization() {
        if (com.google.protobuf.GeneratedMessage
                .alwaysUseFieldBuilders) {
          getCreatedAtFieldBuilder();
        }
      }
      @java.lang.Override
      public Builder clear() {
        super.clear();
        bitField0_ = 0;
        id_ = 0L;
        name_ = "";
        createdAt_ = null;
        if (createdAtBuilder_ != null) {
          createdAtBuilder_.dispose();
          createdAtBuilder_ = null;
        }
        return this;
      }

      @java.lang.Override
      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.internal_static_Manufacturer_descriptor;
      }

      @java.lang.Override
      public com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer getDefaultInstanceForType() {
        return com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer.getDefaultInstance();
      }

      @java.lang.Override
      public com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer build() {
        com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      @java.lang.Override
      public com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer buildPartial() {
        com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer result = new com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer(this);
        if (bitField0_ != 0) { buildPartial0(result); }
        onBuilt();
        return result;
      }

      private void buildPartial0(com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer result) {
        int from_bitField0_ = bitField0_;
        if (((from_bitField0_ & 0x00000001) != 0)) {
          result.id_ = id_;
        }
        if (((from_bitField0_ & 0x00000002) != 0)) {
          result.name_ = name_;
        }
        int to_bitField0_ = 0;
        if (((from_bitField0_ & 0x00000004) != 0)) {
          result.createdAt_ = createdAtBuilder_ == null
              ? createdAt_
              : createdAtBuilder_.build();
          to_bitField0_ |= 0x00000001;
        }
        result.bitField0_ |= to_bitField0_;
      }

      @java.lang.Override
      public Builder mergeFrom(com.google.protobuf.Message other) {
        if (other instanceof com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer) {
          return mergeFrom((com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer)other);
        } else {
          super.mergeFrom(other);
          return this;
        }
      }

      public Builder mergeFrom(com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer other) {
        if (other == com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer.getDefaultInstance()) return this;
        if (other.getId() != 0L) {
          setId(other.getId());
        }
        if (!other.getName().isEmpty()) {
          name_ = other.name_;
          bitField0_ |= 0x00000002;
          onChanged();
        }
        if (other.hasCreatedAt()) {
          mergeCreatedAt(other.getCreatedAt());
        }
        this.mergeUnknownFields(other.getUnknownFields());
        onChanged();
        return this;
      }

      @java.lang.Override
      public final boolean isInitialized() {
        return true;
      }

      @java.lang.Override
      public Builder mergeFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws java.io.IOException {
        if (extensionRegistry == null) {
          throw new java.lang.NullPointerException();
        }
        try {
          boolean done = false;
          while (!done) {
            int tag = input.readTag();
            switch (tag) {
              case 0:
                done = true;
                break;
              case 8: {
                id_ = input.readInt64();
                bitField0_ |= 0x00000001;
                break;
              } // case 8
              case 18: {
                name_ = input.readStringRequireUtf8();
                bitField0_ |= 0x00000002;
                break;
              } // case 18
              case 26: {
                input.readMessage(
                    getCreatedAtFieldBuilder().getBuilder(),
                    extensionRegistry);
                bitField0_ |= 0x00000004;
                break;
              } // case 26
              default: {
                if (!super.parseUnknownField(input, extensionRegistry, tag)) {
                  done = true; // was an endgroup tag
                }
                break;
              } // default:
            } // switch (tag)
          } // while (!done)
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          throw e.unwrapIOException();
        } finally {
          onChanged();
        } // finally
        return this;
      }
      private int bitField0_;

      private long id_ ;
      /**
       * <code>int64 id = 1;</code>
       * @return The id.
       */
      @java.lang.Override
      public long getId() {
        return id_;
      }
      /**
       * <code>int64 id = 1;</code>
       * @param value The id to set.
       * @return This builder for chaining.
       */
      public Builder setId(long value) {

        id_ = value;
        bitField0_ |= 0x00000001;
        onChanged();
        return this;
      }
      /**
       * <code>int64 id = 1;</code>
       * @return This builder for chaining.
       */
      public Builder clearId() {
        bitField0_ = (bitField0_ & ~0x00000001);
        id_ = 0L;
        onChanged();
        return this;
      }

      private java.lang.Object name_ = "";
      /**
       * <code>string name = 2;</code>
       * @return The name.
       */
      public java.lang.String getName() {
        java.lang.Object ref = name_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          name_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string name = 2;</code>
       * @return The bytes for name.
       */
      public com.google.protobuf.ByteString
          getNameBytes() {
        java.lang.Object ref = name_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          name_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string name = 2;</code>
       * @param value The name to set.
       * @return This builder for chaining.
       */
      public Builder setName(
          java.lang.String value) {
        if (value == null) { throw new NullPointerException(); }
        name_ = value;
        bitField0_ |= 0x00000002;
        onChanged();
        return this;
      }
      /**
       * <code>string name = 2;</code>
       * @return This builder for chaining.
       */
      public Builder clearName() {
        name_ = getDefaultInstance().getName();
        bitField0_ = (bitField0_ & ~0x00000002);
        onChanged();
        return this;
      }
      /**
       * <code>string name = 2;</code>
       * @param value The bytes for name to set.
       * @return This builder for chaining.
       */
      public Builder setNameBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) { throw new NullPointerException(); }
        checkByteStringIsUtf8(value);
        name_ = value;
        bitField0_ |= 0x00000002;
        onChanged();
        return this;
      }

      private com.google.protobuf.Timestamp createdAt_;
      private com.google.protobuf.SingleFieldBuilder<
          com.google.protobuf.Timestamp, com.google.protobuf.Timestamp.Builder, com.google.protobuf.TimestampOrBuilder> createdAtBuilder_;
      /**
       * <code>.google.protobuf.Timestamp created_at = 3;</code>
       * @return Whether the createdAt field is set.
       */
      public boolean hasCreatedAt() {
        return ((bitField0_ & 0x00000004) != 0);
      }
      /**
       * <code>.google.protobuf.Timestamp created_at = 3;</code>
       * @return The createdAt.
       */
      public com.google.protobuf.Timestamp getCreatedAt() {
        if (createdAtBuilder_ == null) {
          return createdAt_ == null ? com.google.protobuf.Timestamp.getDefaultInstance() : createdAt_;
        } else {
          return createdAtBuilder_.getMessage();
        }
      }
      /**
       * <code>.google.protobuf.Timestamp created_at = 3;</code>
       */
      public Builder setCreatedAt(com.google.protobuf.Timestamp value) {
        if (createdAtBuilder_ == null) {
          if (value == null) {
            throw new NullPointerException();
          }
          createdAt_ = value;
        } else {
          createdAtBuilder_.setMessage(value);
        }
        bitField0_ |= 0x00000004;
        onChanged();
        return this;
      }
      /**
       * <code>.google.protobuf.Timestamp created_at = 3;</code>
       */
      public Builder setCreatedAt(
          com.google.protobuf.Timestamp.Builder builderForValue) {
        if (createdAtBuilder_ == null) {
          createdAt_ = builderForValue.build();
        } else {
          createdAtBuilder_.setMessage(builderForValue.build());
        }
        bitField0_ |= 0x00000004;
        onChanged();
        return this;
      }
      /**
       * <code>.google.protobuf.Timestamp created_at = 3;</code>
       */
      public Builder mergeCreatedAt(com.google.protobuf.Timestamp value) {
        if (createdAtBuilder_ == null) {
          if (((bitField0_ & 0x00000004) != 0) &&
            createdAt_ != null &&
            createdAt_ != com.google.protobuf.Timestamp.getDefaultInstance()) {
            getCreatedAtBuilder().mergeFrom(value);
          } else {
            createdAt_ = value;
          }
        } else {
          createdAtBuilder_.mergeFrom(value);
        }
        if (createdAt_ != null) {
          bitField0_ |= 0x00000004;
          onChanged();
        }
        return this;
      }
      /**
       * <code>.google.protobuf.Timestamp created_at = 3;</code>
       */
      public Builder clearCreatedAt() {
        bitField0_ = (bitField0_ & ~0x00000004);
        createdAt_ = null;
        if (createdAtBuilder_ != null) {
          createdAtBuilder_.dispose();
          createdAtBuilder_ = null;
        }
        onChanged();
        return this;
      }
      /**
       * <code>.google.protobuf.Timestamp created_at = 3;</code>
       */
      public com.google.protobuf.Timestamp.Builder getCreatedAtBuilder() {
        bitField0_ |= 0x00000004;
        onChanged();
        return getCreatedAtFieldBuilder().getBuilder();
      }
      /**
       * <code>.google.protobuf.Timestamp created_at = 3;</code>
       */
      public com.google.protobuf.TimestampOrBuilder getCreatedAtOrBuilder() {
        if (createdAtBuilder_ != null) {
          return createdAtBuilder_.getMessageOrBuilder();
        } else {
          return createdAt_ == null ?
              com.google.protobuf.Timestamp.getDefaultInstance() : createdAt_;
        }
      }
      /**
       * <code>.google.protobuf.Timestamp created_at = 3;</code>
       */
      private com.google.protobuf.SingleFieldBuilder<
          com.google.protobuf.Timestamp, com.google.protobuf.Timestamp.Builder, com.google.protobuf.TimestampOrBuilder> 
          getCreatedAtFieldBuilder() {
        if (createdAtBuilder_ == null) {
          createdAtBuilder_ = new com.google.protobuf.SingleFieldBuilder<
              com.google.protobuf.Timestamp, com.google.protobuf.Timestamp.Builder, com.google.protobuf.TimestampOrBuilder>(
                  getCreatedAt(),
                  getParentForChildren(),
                  isClean());
          createdAt_ = null;
        }
        return createdAtBuilder_;
      }

      // @@protoc_insertion_point(builder_scope:Manufacturer)
    }

    // @@protoc_insertion_point(class_scope:Manufacturer)
    private static final com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer();
    }

    public static com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<Manufacturer>
        PARSER = new com.google.protobuf.AbstractParser<Manufacturer>() {
      @java.lang.Override
      public Manufacturer parsePartialFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws com.google.protobuf.InvalidProtocolBufferException {
        Builder builder = newBuilder();
        try {
          builder.mergeFrom(input, extensionRegistry);
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          throw e.setUnfinishedMessage(builder.buildPartial());
        } catch (com.google.protobuf.UninitializedMessageException e) {
          throw e.asInvalidProtocolBufferException().setUnfinishedMessage(builder.buildPartial());
        } catch (java.io.IOException e) {
          throw new com.google.protobuf.InvalidProtocolBufferException(e)
              .setUnfinishedMessage(builder.buildPartial());
        }
        return builder.buildPartial();
      }
    };

    public static com.google.protobuf.Parser<Manufacturer> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<Manufacturer> getParserForType() {
      return PARSER;
    }

    @java.lang.Override
    public com.nyrta1.SkyTicket_AdminService.models.ManufacturerEntity.Manufacturer getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_Manufacturer_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_Manufacturer_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\030proto/manufacturer.proto\032\037google/proto" +
      "buf/timestamp.proto\"X\n\014Manufacturer\022\n\n\002i" +
      "d\030\001 \001(\003\022\014\n\004name\030\002 \001(\t\022.\n\ncreated_at\030\003 \001(" +
      "\0132\032.google.protobuf.TimestampB>\n(com.nyr" +
      "ta1.SkyTicket_AdminService.modelsB\022Manuf" +
      "acturerEntityb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.google.protobuf.TimestampProto.getDescriptor(),
        });
    internal_static_Manufacturer_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_Manufacturer_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_Manufacturer_descriptor,
        new java.lang.String[] { "Id", "Name", "CreatedAt", });
    descriptor.resolveAllFeaturesImmutable();
    com.google.protobuf.TimestampProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
